package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/redis/go-redis/v9"
)

func main() {
	logger := watermill.NewStdLogger(false, false)

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	sub, err := redisstream.NewSubscriber(redisstream.SubscriberConfig{
		Client: rdb,
	}, logger)
	if err != nil {
		panic(err)
	}

	// pub, err := redisstream.NewPublisher(redisstream.PublisherConfig{
	// 	Client: rdb,
	// }, logger)
	// if err != nil {
	// 	panic(err)
	// }

	router.AddNoPublisherHandler(
		"fahrenheit-reader",
		"temperature-fahrenheit",
		sub,
		func(msg *message.Message) error {
			fmt.Printf("Temperature read: %s\n", msg.Payload)
			return nil
		},
	)

	err = router.Run(context.Background())
	if err != nil {
		panic(err)
	}
}
