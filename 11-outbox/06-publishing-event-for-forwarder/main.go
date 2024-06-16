package main

import (
	"database/sql"

	"github.com/ThreeDotsLabs/watermill"
	watermillSql "github.com/ThreeDotsLabs/watermill-sql/v3/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/components/forwarder"
	"github.com/ThreeDotsLabs/watermill/message"

	_ "github.com/lib/pq"
)

var outboxTopic = "events_to_forward"

func PublishInTx(
	msg *message.Message,
	tx *sql.Tx,
	logger watermill.LoggerAdapter,
) error {
	// your code goes here
	sqlPublisher, err := watermillSql.NewPublisher(
		tx,
		watermillSql.PublisherConfig{
			SchemaAdapter: watermillSql.DefaultPostgreSQLSchema{},
		},
		logger,
	)
	if err != nil {
		return err
	}
	publisher := forwarder.NewPublisher(sqlPublisher, forwarder.PublisherConfig{
		ForwarderTopic: outboxTopic,
	})
	return publisher.Publish("ItemAddedToCart", msg)
}
