package options

import "time"

type KafkaArgs struct {
	Address       string
	Network       string
	Topic         string
	GroupId       string
	PartitionId   int
	Timeout       time.Duration
	WriteDeadline time.Duration
}
