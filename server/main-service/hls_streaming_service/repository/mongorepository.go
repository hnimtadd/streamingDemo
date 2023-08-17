package repositories

import (
	"context"
	"errors"
	"log"
	"main-service/main-service/config"
	"main-service/main-service/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type hlsStreamRepositoryMongo struct {
	config config.MongoConfig
	db     *mongo.Database
}

func CreateMongoRepo(config config.MongoConfig) HlsStreamingRepo {
	repo := &hlsStreamRepositoryMongo{config: config}
	if err := repo.initDb(); err != nil {
		log.Fatalf("error: %v", err)
	}
	return repo
}

func (s *hlsStreamRepositoryMongo) initDb() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOpts := options.Client().ApplyURI(s.config.DbSource)
	// SetAuth(
	// options.Credential{
	// 	AuthSource:    s.config.DbAuthSource,
	// 	AuthMechanism: s.config.DbAuthMechanism,
	// 	Username:      s.config.DbUsername,
	// 	Password:      s.config.DbPassword,
	// },
	// ).SetTLSConfig(nil)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	s.db = client.Database(s.config.DbDatabase)
	return nil

}

func (s *hlsStreamRepositoryMongo) CheckThenInsertCameras(cam *entities.Camera) (bool, error) {
	// false indicate this camera is new camera, and db repository inserted it into db, vice versa
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	filter := bson.D{primitive.E{Key: "sourceurl", Value: cam.SourceUrl}}
	res := s.db.Collection("cameras").FindOneAndReplace(ctx, filter, cam)
	if err := res.Err(); err != nil {
		if err != mongo.ErrNoDocuments {
			return false, err
		}
		if err := s.InsertCameras(cam); err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (s *hlsStreamRepositoryMongo) InsertCameras(cam *entities.Camera) error {
	if cam.SourceUrl == "" {
		return errors.New("SourceUrl must not null")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := s.db.Collection("cameras").InsertOne(ctx, cam)
	if err != nil {
		return err
	}
	return nil
}

func (s *hlsStreamRepositoryMongo) GetCameras() ([]CameraSource, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	cur, err := s.db.Collection("cameras").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	res := []CameraSource{}
	for cur.Next(ctx) {
		var cam CameraSource
		if err := cur.Decode(&cam); err != nil {
			return nil, err
		}
		res = append(res, cam)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return res, nil

}
