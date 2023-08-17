package services

import (
	"context"
	"encoding/json"
	"log"
	"main-service/main-service/config"
	"main-service/main-service/entities"
	message "main-service/main-service/entities/message"
	"sync"

	repositories "main-service/main-service/hls_streaming_service/repository"

	"github.com/segmentio/kafka-go"
)

type HlsStreamingService interface {
	SubscribeEventResponseMessages() error
	GetCameras() ([]repositories.CameraSource, error)

	InitBackground()
}

type hlsStreamingService struct {
	repo   repositories.HlsStreamingRepo
	config config.KafkaConfig
}

func CreateService(
	repo repositories.HlsStreamingRepo,
	config config.KafkaConfig,
) HlsStreamingService {
	return &hlsStreamingService{
		repo:   repo,
		config: config,
	}
}

func (s *hlsStreamingService) InitBackground() {
	// Read for message in kafka queue, process message, suppose that there only camera management add request
	s.SubscribeEventResponseMessages()
}

func (s *hlsStreamingService) SubscribeEventResponseMessages() error {
	// TODO: Subsribe to the CAMERA_MANAGEMENT kafka topic and inserting it to the database

	// Hardcode part, should move to service config
	// Open connection to kafka server
	config := kafka.ReaderConfig{
		Brokers: []string{s.config.Address},
		Topic:   s.config.Topic,
		GroupID: s.config.GroupId,
	}

	r := kafka.NewReader(config)
	defer r.Close()
	var wg sync.WaitGroup

	ctx := context.Background()
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("[Service] Listening on queue")
		for {
			d, err := r.ReadMessage(ctx)
			if err != nil {
				log.Printf("error: %v", err)
				break
			}
			log.Println("[Service] Received message")
			if err := s.InsertCameras(d); err != nil {
				log.Fatalf("error: %v", err)
			}
		}
		log.Println("[Service] Complete listening on queue")
	}()
	wg.Wait()
	return nil
}

func (s *hlsStreamingService) InsertCameras(d kafka.Message) error {
	var msgs []message.CameraManagementMessage
	if err := json.Unmarshal(d.Value, &msgs); err != nil {
		return err
	}

	// TODO: representer layer, transform msg => camera entity to forward to repo
	// var cams []entities.Camera
	for _, msg := range msgs {
		cam := entities.Camera{
			SourceUrl:                  msg.SourceUrl,
			CameraHlsStreamingEndpoint: msg.CameraHlsStreamingEndpoint,
		}

		if contain, err := s.repo.CheckThenInsertCameras(&cam); err != nil {
			// Error while inserting cameras, should continue insert and print out which camera info isn't insert success
			log.Printf("[ERROR]: %v --- camera : %v", err, cam)
		} else if !contain {
			log.Printf("[ADDED] cam: %v", cam)
		} else {
			log.Printf("[UPDATED] cam: %v", cam)
		}
		// cams = append(cams, cam)
	}
	return nil
}

func (s *hlsStreamingService) GetCameras() ([]repositories.CameraSource, error) {
	cams, err := s.repo.GetCameras()
	if err != nil {
		return nil, err
	}

	return cams, nil
}
