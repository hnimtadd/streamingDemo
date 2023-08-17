package config

type MongoConfig struct {
	DbSource        string `json:"db_source"`
	DbUsername      string `json:"db_username"`
	DbPassword      string `json:"db_password"`
	DbAuthSource    string `json:"db_authsource"`
	DbAuthMechanism string `json:"db_authmechanism"`
	DbDatabase      string `json:"db_database"`
}

func NewMongoConfig() MongoConfig {
	return MongoConfig{
		// DbSource:     "mongodb://video-analytics-mongodb-hls-streaming:27017/",
		// DbSource:        "mongodb://vinai_User:vinai_Password@localhost:27017/video_analytis?parseTime=true&authSource=video_analytics",
		DbSource:        "mongodb://vinai_User:vinai_Password@video-analytics-mongodb-hls-streaming:27017/video_analytis?parseTime=true&authSource=video_analytics",
		DbDatabase:      "video_analytics",
		DbUsername:      "vinai_User",
		DbAuthSource:    "video_analytics",
		DbPassword:      "vinai_Password",
		DbAuthMechanism: "SCRAM-SHA-1",
	}
}
