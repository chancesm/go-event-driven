package main

import (
	"os"
	"strconv"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/redis/go-redis/v9"
)

func main() {
	logger := watermill.NewStdLogger(false, false)

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	publisher, err := redisstream.NewPublisher(redisstream.PublisherConfig{
		Client: rdb,
	}, logger)
	if err != nil {
		panic(err)
	}

	err = publisher.Publish("progress", progressMessage(50), progressMessage(100))
	if err != nil {
		panic(err)
	}
}
func progressMessage(value int) *message.Message {
	return message.NewMessage(watermill.NewUUID(), []byte(strconv.Itoa(value)))
}
