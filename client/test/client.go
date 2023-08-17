package main

import (
	"cameraClient/client"
	"cameraClient/config"
	"log"
)

func main() {
	clientInfo := config.NewClientConfig()
	cameraClient, err := client.NewClient(clientInfo)

	// Kafka enable for publishing camera info from client to serve
	// opts := options.KafkaArgs{
	// 	Network:       publishConfig.Network,
	// 	Address:       publishConfig.Address,
	// 	Topic:         publishConfig.Topic,
	// 	GroupId:       publishConfig.GroupId,
	// 	PartitionId:   publishConfig.PartitionId,
	// 	Timeout:       time.Second * 5,
	// 	WriteDeadline: time.Second * 10,
	// }
	// cameraClient.WithKafka(opts)
	// if err != nil {
	// 	log.Fatalf("err: %v", err)
	// }

	// Get possible network interfaces in network
	interfaces, err := cameraClient.GetCameraInterface()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Start scan onvif camera in posibble interfaces
	camerasInfo, err := cameraClient.ScanCamera(interfaces...)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if len(camerasInfo) == 0 {
		log.Fatalf("There's no cameras streaming in this network")
	}

	log.Println(camerasInfo)

	// Get camera entities, which can run streamTo function to stream to server
	camera, err := cameraClient.GetCamera(camerasInfo[0])
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	CameraRTMPStreamingEndpoint := "rtmp://171.244.62.138:1935/live/sample"

	// Register camera information locally, if this is new camera, return contain, err
	_, err = camera.RegisterEndPoint(CameraRTMPStreamingEndpoint)
	// contain, err := camera.RegisterEndPoint(CameraHlsStreamingEndpoint)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// if camera is not registered locally, maybe this camera is not register at server
	// if !contain {
	// 	log.Println("Endpoint not published, trying to publish")
	// 	if err := cameraClient.PublishCameraInfoToServer(camera); err != nil {
	// 		log.Fatalf("error: %v", err)
	// 	}
	// }

	if err := camera.StreamTo(CameraRTMPStreamingEndpoint); err != nil {
		log.Fatalf("error: %v", err)
	}
}
