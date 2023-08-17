package client

import (
	"cameraClient/config"
	"cameraClient/entities"
	"cameraClient/options"
	"cameraClient/utils"
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"
)

type CameraClient interface {
	/*
		Scan possible network interfaces in current network using net.interface
			- output: []string: name of network interfaces
	*/
	GetCameraInterface() ([]string, error)

	/*
		Scan onvif camera running in interfaces
			-input: []string: name of netwok interfaces to scan
			-output: []entities.Camera, and err if occurs
	*/
	ScanCamera(...string) ([]entities.Camera, error)

	GetCamerasInfo() []entities.Camera
	GetCamera(entities.Camera) (Camera, error)
	GetHardwareInfo() *entities.HardwareInfo
	PublishCameraInfoToServer(cam Camera) error

	/*
		Enable Kafka agent in client, using for publishing information to kafka server
	*/
	WithKafka(opts options.KafkaArgs) *cameraClient
}

type cameraClient struct {
	Cameras      []entities.Camera
	HardwareInfo entities.HardwareInfo
	Publisher    *kafkaPublisher
	ClientInfo   entities.Client
}

func NewClient(clientInfo config.ClientConfig) (CameraClient, error) {
	client := &cameraClient{}
	if err := client.initClient(); err != nil {
		return client, err
	}

	client.ClientInfo = entities.Client{}
	return client, nil
}

func (s *cameraClient) initClient() error {
	hardwareInfo, err := utils.GetHardwareInfo()
	if err != nil {
		return err
	}
	s.HardwareInfo = hardwareInfo
	return nil
}

func (s *cameraClient) GetCamerasInfo() []entities.Camera {
	return s.Cameras
}

func (s *cameraClient) GetCamera(info entities.Camera) (Camera, error) {
	camera, err := NewCamera(info)
	if err != nil {
		return nil, err
	}
	return camera, nil
}

func (s *cameraClient) GetHardwareInfo() *entities.HardwareInfo {
	return &s.HardwareInfo
}

func (s *cameraClient) PublishCameraInfoToServer(cam Camera) error {
	if s.Publisher == nil {
		log.Fatalf("error: publisher not initialize")
	}
	camInfo := cam.GetInfo()
	endpoints := cam.GetEndpoints()
	if len(endpoints) == 0 {
		return errors.New("endpoints must exists")
	}

	msg := entities.CameraManagementMessage{
		Cameras: []entities.Camera{camInfo},
		Client:  s.ClientInfo,
	}
	body, err := json.Marshal(&msg)
	if err != nil {
		return err
	}
	req := PublishRequest{
		Body: body,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := s.Publisher.kafkaPublish(ctx, req); err != nil {
		return err
	}
	log.Println("[client] published camera info to kafka")
	return nil
}
