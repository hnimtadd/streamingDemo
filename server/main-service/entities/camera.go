package entities

type Camera struct {
	EntityBase
	CameraId                   string `json:"camera_id,omitempty" db:"camera_id,omitempty"`
	CameraHlsStreamingEndpoint string `json:"camera_hls_streaming_endpoint,omitempty" db:"camera_hls_streaming_endpoint,omitempty"`
	SourceUrl                  string `json:"source_url,omitempty" db:"source_url,omitempty"`
}
