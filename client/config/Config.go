package config

type CameraConfig struct {
	RtspUrl string `json:"rtsp_url"`
}

func NewCameraConfig() CameraConfig {
	return CameraConfig{
		RtspUrl: "",
	}
}

type PublishConfig struct {
	Network     string `json:"network"`
	Topic       string `json:"topic"`
	Address     string `json:"address"`
	GroupId     string `json:"group_id"`
	PartitionId int    `json:"partition_id"`
}

func NewPublishConfig() PublishConfig {
	return PublishConfig{
		Topic:   "CAMERA_MANAGEMENT",
		Network: "tcp",
		Address: "localhost:29092",
		// Address:     "kafka:9092",
		GroupId:     "serviceworker",
		PartitionId: 0,
	}
}

type ClientConfig struct {
	Id       string `json:"id,omiempty"`
	Location string `json:"location,omiempty"`
}

func NewClientConfig() ClientConfig {
	return ClientConfig{
		Id:       "sampleId",
		Location: "sampleLocation",
	}
}
