package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"tickets/api"
	"tickets/message"
	"tickets/service"

	_ "github.com/lib/pq"

	"github.com/ThreeDotsLabs/go-event-driven/common/clients"
	"github.com/ThreeDotsLabs/go-event-driven/common/log"
	"github.com/jmoiron/sqlx"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	apiClients, err := clients.NewClients(
		os.Getenv("GATEWAY_ADDR"),
		func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Correlation-ID", log.CorrelationIDFromContext(ctx))
			return nil
		},
	)
	if err != nil {
		panic(err)
	}

	db, err := sqlx.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	redisClient := message.NewRedisClient(os.Getenv("REDIS_ADDR"))
	defer redisClient.Close()

	deadNationAPI := api.NewDeadNationClient(apiClients)
	spreadsheetsService := api.NewSpreadsheetsAPIClient(apiClients)
	receiptsService := api.NewReceiptsServiceClient(apiClients)
	filesAPI := api.NewFilesApiClient(apiClients)

	err = service.New(
		db,
		redisClient,
		deadNationAPI,
		spreadsheetsService,
		receiptsService,
		filesAPI,
	).Run(ctx)
	if err != nil {
		panic(err)
	}
}
