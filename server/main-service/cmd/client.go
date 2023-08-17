package main

import (
	"context"
	"encoding/json"
	"log"
	entities "main-service/main-service/entities/message"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Test File
	message := []entities.CameraManagementMessage{
		{
			CameraHlsStreamingEndpoint: "randomtext1",
			SourceUrl:                  "randomtext1",
		},
		{
			CameraHlsStreamingEndpoint: "randomtext2",
			SourceUrl:                  "randomtext2",
		},
	}

	topic := "CAMERA_MANAGEMENT"
	partition := 0
	network := "tcp"
	address := "localhost:39092"
	conn, err := kafka.DialLeader(context.Background(), network, address, topic, partition)
	log.Println(address)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	body, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	kafkamessage := kafka.Message{
		Value: body,
	}
	_, err = conn.WriteMessages(kafkamessage)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	log.Printf("Published message")

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
