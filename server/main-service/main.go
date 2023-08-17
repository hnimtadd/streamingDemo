package main

import (
	"main-service/main-service/config"
	hls_api_handler "main-service/main-service/hls_streaming_service/handlers"
	hls_repo "main-service/main-service/hls_streaming_service/repository"
	hls_sv "main-service/main-service/hls_streaming_service/service"
	"sync"
)

func main() {
	// HLS streaming service
	// repoConfig := config.NewMysqlConfig()
	// hls_streaming_repository := hls_repo.CreateRepo(repoConfig)
	repoConfig := config.NewMongoConfig()
	hls_streaming_repository := hls_repo.CreateMongoRepo(repoConfig)

	kafkaConfig := config.NewKafkaConfig()
	hls_streaming_service := hls_sv.CreateService(hls_streaming_repository, kafkaConfig)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		hls_streaming_service.InitBackground()
	}()

	hls_streaming_api_handler := hls_api_handler.CreateApiHandler(hls_streaming_service)

	hls_streaming_api_handler.Run(":10077")
	wg.Wait()
}
