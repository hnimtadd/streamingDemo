package entities

type CameraManagementMessage struct {
	CameraHlsStreamingEndpoint string `json:"camera_hls_streaming_endpoint,omitempty" db:"camera_hls_streaming_endpoint,omitempty"`
	SourceUrl                  string `json:"source_url,omitempty" db:"source_url,omitempty"`
}
