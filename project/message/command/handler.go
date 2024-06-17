package command

import (
	"context"
	"tickets/entities"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/google/uuid"
)

type Handler struct {
	commandBus *cqrs.CommandBus
}

func NewHandler(
	commandBus *cqrs.CommandBus,
) Handler {
	if commandBus == nil {
		panic("missing commandBus")
	}

	return Handler{
		commandBus: commandBus,
	}
}

type SpreadsheetsAPI interface {
	AppendRow(ctx context.Context, sheetName string, row []string) error
}

type ReceiptsService interface {
	IssueReceipt(ctx context.Context, request entities.IssueReceiptRequest) (entities.IssueReceiptResponse, error)
}

type FilesAPI interface {
	UploadFile(ctx context.Context, fileID string, fileContent string) error
}

type TicketsRepository interface {
	Add(ctx context.Context, ticket entities.Ticket) error
	Remove(ctx context.Context, ticketID string) error
}

type ShowsRepository interface {
	ShowByID(ctx context.Context, showID uuid.UUID) (entities.Show, error)
}

type DeadNationAPI interface {
	BookInDeadNation(ctx context.Context, request entities.DeadNationBooking) error
}
