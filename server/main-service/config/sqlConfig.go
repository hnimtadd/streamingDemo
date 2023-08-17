package config

type MysqlConfig struct {
	DriverName string `json:"db_source"`
	Dsn        string `json:"db_dsn"`
}

func NewMysqlConfig() MysqlConfig {
	return MysqlConfig{
		DriverName: "mysql",
		Dsn:        "vinai_User:vinai_Password@tcp(video-analytics-mysql-hls-streaming:3306)/video_analytics?parseTime=true",
		// Dsn: "vinai_User:vinai_Password@tcp(localhost:3306)/video_analytics?parseTime=true",
	}
}
