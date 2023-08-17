package entities

type Camera struct {
	SourceUrls []SourceUrl `json:"source_urls,omiempty" db:"source_urls,omiempty"`
	Comment    string      `json:"comment,omiempty" db:"comment,omiempty"`
}

type SourceUrl struct {
	SourceType string `json:"type,omiempty" db:"type,omiempty"`
	SourceUrl  string `json:"url,omiempty" db:"url,omiempty"`
}

type CameraManagementMessage struct {
	Client
	Cameras []Camera `json:"cameras,omiempty" db:"cameras,omiemtpy"`
}
