package client

import (
	"cameraClient/options"
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type PublishRequest struct {
	Body []byte
}

// type Publisher interface {
// 	Publish(context.Context, PublishRequest) error
// }

type kafkaPublisher struct {
	conn *kafka.Conn
	args options.KafkaArgs
}

func newKafkaPublisher(args options.KafkaArgs) (kafkaPublisher, error) {
	publisher := kafkaPublisher{
		args: args,
	}
	if err := publisher.initPublisher(); err != nil {
		return publisher, err
	}
	return publisher, nil
}

func (s *kafkaPublisher) initPublisher() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.args.Timeout)
	defer cancel()
	conn, err := kafka.DialLeader(ctx, s.args.Network, s.args.Address, s.args.Topic, s.args.PartitionId)
	if err != nil {
		return err
	}
	s.conn = conn
	return nil
}

func (s *kafkaPublisher) kafkaPublish(ctx context.Context, req PublishRequest) error {
	kafkaMessage := kafka.Message{
		Value: req.Body,
	}
	if _, err := s.conn.WriteMessages(kafkaMessage); err != nil {
		return err
	}
	log.Println("Published")
	return nil
}

func (s *cameraClient) WithKafka(opts options.KafkaArgs) *cameraClient {
	publisher, err := newKafkaPublisher(opts)
	if err != nil {
		log.Fatalf("Cann't init kafka publisher, error: %v", err)
	}
	s.Publisher = &publisher
	return s
}
