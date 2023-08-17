package config

type KafkaConfig struct {
	Topic   string `json:"topic"`
	Address string `json:"address"`
	GroupId string `json:"group_id"`
}

func NewKafkaConfig() KafkaConfig {
	return KafkaConfig{
		Topic: "CAMERA_MANAGEMENT",
		// Address: "localhost:29092",
		Address: "kafka:19092",
		GroupId: "serviceworker",
	}
}
