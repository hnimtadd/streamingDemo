package api_handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/segmentio/kafka-go"

	"main-service/main-service/config"
	"main-service/main-service/entities/message"
	services "main-service/main-service/hls_streaming_service/service"
)

type HlsStreamingApiHandler interface {
	HandleGetCameras(w http.ResponseWriter, r *http.Request)

	Run(port string)
}

type hlsStreamingApiHandler struct {
	hlsStreamingService services.HlsStreamingService
}

func CreateApiHandler(
	hlsStreamingService services.HlsStreamingService,
) HlsStreamingApiHandler {
	return &hlsStreamingApiHandler{
		hlsStreamingService: hlsStreamingService,
	}
}

func (h *hlsStreamingApiHandler) Run(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/event/get-cameras", h.HandleGetCameras).Methods("GET")
	r.HandleFunc("/api/v1/event/publish", h.HandlePublishMessage).Methods("POST")

	handler := cors.Default().Handler(r)

	log.Println("Listening on port: ", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (h *hlsStreamingApiHandler) HandleGetCameras(w http.ResponseWriter, r *http.Request) {
	// TODO: Handle GET API request /api/v1/event/get-cameras

	cams, err := h.hlsStreamingService.GetCameras()
	// check if error occurs while request for cameras source
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cams)
	return
}

func (h *hlsStreamingApiHandler) HandlePublishMessage(w http.ResponseWriter, r *http.Request) {
	// //TODO: hardcode part,
	// client must include some specific information, at least userid
	var message []entities.CameraManagementMessage
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf("error: %v", err)
	}

	config := config.NewKafkaConfig()
	partition := 0
	network := "tcp"
	conn, err := kafka.DialLeader(context.Background(), network, config.Address, config.Topic, partition)
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
