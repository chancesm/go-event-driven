package main

import (
	"database/sql"

	"github.com/ThreeDotsLabs/watermill"
	watermillSql "github.com/ThreeDotsLabs/watermill-sql/v3/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/message"
	_ "github.com/lib/pq"
)

func PublishInTx(
	message *message.Message,
	tx *sql.Tx,
	logger watermill.LoggerAdapter,
) error {
	// your code goes here
	publisher, err := watermillSql.NewPublisher(
		tx,
		watermillSql.PublisherConfig{
			SchemaAdapter: watermillSql.DefaultPostgreSQLSchema{},
		},
		logger,
	)
	if err != nil {
		panic("could not init publisher")
	}
	publisher.Publish("ItemAddedToCart", message)
	return nil
}
