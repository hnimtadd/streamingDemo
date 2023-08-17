package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"main-service/main-service/config"
	dbqueries "main-service/main-service/database_queries"
	"main-service/main-service/entities"
)

type CameraSource struct {
	CameraId                   string `json:"camera_id,omitempty" db:"id,omitempty" bson:"_id"`
	CameraHlsStreamingEndpoint string `json:"camera_hls_streaming_endpoint,omitempty" db:"camera_hls_streaming_endpoint,omitempty"`
	SourceUrl                  string `json:"source_url,omitempty" db:"source_url,omitempty"`
}

type HlsStreamingRepo interface {
	// bool indicate that this camera exists in db or not
	CheckThenInsertCameras(cam *entities.Camera) (bool, error)
	InsertCameras(cam *entities.Camera) error
	GetCameras() ([]CameraSource, error)
}

type hlsStreamingRepo struct {
	db *sql.DB
}

func CreateRepo(config config.MysqlConfig) HlsStreamingRepo {
	repo := &hlsStreamingRepo{}
	if err := repo.initRepo(config); err != nil {
		log.Fatalf("error while init db: %v", err)
	}
	return repo
}

func (s *hlsStreamingRepo) initRepo(config config.MysqlConfig) error {
	db, err := sql.Open(config.DriverName, config.Dsn)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *hlsStreamingRepo) InsertCameras(cam *entities.Camera) error {
	// TODO: Insert the cameras to the database

	// CREATE_CAMERA = ` INSERT INTO video_analytics.cameras(camera_id, camera_hls_streaming_endpoint, source_url)
	// VALUES (?, ?, ?)
	if _, err := s.db.Exec(dbqueries.CREATE_CAMERA, cam.CameraHlsStreamingEndpoint, cam.SourceUrl); err != nil {
		return err
	}
	return nil
}

func (s *hlsStreamingRepo) GetCameras() ([]CameraSource, error) {
	// TODO: Return the cameras from the database

	// GET_CAMERA = ` SELECT camera_id, camera_hls_streaming_endpoint, source_url
	// 	FROM video_analytics.cameras WHERE 1
	// `
	rows, err := s.db.Query(dbqueries.GET_CAMERA)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cams := []CameraSource{}
	for rows.Next() {
		var cam CameraSource
		if err := rows.Scan(&cam.CameraId, &cam.CameraHlsStreamingEndpoint, &cam.SourceUrl); err != nil {
			return nil, err
		}
		cams = append(cams, cam)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cams, nil
}

func (s *hlsStreamingRepo) CheckThenInsertCameras(cam *entities.Camera) (bool, error) {
	return false, nil
}
