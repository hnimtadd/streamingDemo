package client

import (
	"cameraClient/entities"
	"cameraClient/utils"
	"errors"
	"log"
	"os"
	"path/filepath"
)

type Camera interface {
	/*input: endpointURL, url to stream to
	  output: error if occur, else nil */
	StreamTo(endpointURL string) error
	// input: register camera source url locally
	// output: contain (bool) => true if camera already registered, else false
	RegisterEndPoint(url string) (contain bool, err error)
	GetEndpoints() []string
	GetInfo() entities.Camera
}

type camera struct {
	info      entities.Camera
	endpoints []string
}

func NewCamera(info entities.Camera) (Camera, error) {
	camera := &camera{
		info: info,
	}
	if err := camera.checkCamera(); err != nil {
		return nil, err
	}
	if err := camera.LoadSavedEndpoint(); err != nil {
		return nil, err
	}
	return camera, nil
}

func (s *camera) GetInfo() entities.Camera {
	return s.info
}

func (s *camera) LoadSavedEndpoint() error {
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(curDir, "data", "endpoints.txt")
	endpoints, err := utils.Readlines(path)
	if err != nil {
		return err
	}
	log.Println("[camera] Loaded endpoints from file: ", endpoints)
	s.endpoints = endpoints
	return nil
}

func (s *camera) SaveEndpoint(endpoint string) error {
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(curDir, "data", "endpoints.txt")
	if err := utils.Writeline(path, endpoint); err != nil {
		return err
	}
	log.Println("[camera] Saved endpoint to file: ", endpoint)
	return nil
}

func (s *camera) RegisterEndPoint(url string) (contain bool, err error) {
	for _, endpoint := range s.endpoints {
		if endpoint == url {
			contain = true
			return
		}
	}
	contain = false
	err = s.SaveEndpoint(url)
	if err != nil {
		return
	}
	s.endpoints = append(s.endpoints, url)
	err = nil
	return
}

func (s *camera) GetEndpoints() []string {
	return s.endpoints
}

func (s *camera) checkCamera() error {
	//TODO: some stuffs todo to check camera is streaming with rtsp
	return nil
}

func (s *camera) StreamTo(endpoint string) error {
	inputArgs := map[string]interface{}{
		"re":     "",
		"an":     "",
		"sn":     "",
		"fflags": "nobuffer",
		"flags":  "low_delay",
	}
	outputArgs := map[string]interface{}{
		"c:v": "copy",
		"an":  "",
		"sn":  "",
		"f":   "flv",
		"y":   "",
	}
	log.Println(s)
	var rtspSource string
	for _, source := range s.info.SourceUrls {
		if source.SourceType == "rtsp" {
			rtspSource = source.SourceUrl
		}
	}
	if rtspSource == "" {
		return errors.New("Camera don't have rtsp source")

	}
	if err := utils.StreamFFmpeg(rtspSource, endpoint, inputArgs, outputArgs); err != nil {
		return err
	}
	return nil
}
