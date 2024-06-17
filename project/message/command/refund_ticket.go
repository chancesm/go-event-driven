package command

import (
	"context"
	"tickets/entities"
)

func (h Handler) RefundTicket(ctx context.Context, event *entities.RefundTicket) error {
	// log.FromContext(ctx).Info("Issuing receipt")

	// request := entities.PaymentRefundRequest{
	// 	TicketID:       event.TicketID,
	// 	Price:          event.Price,
	// 	IdempotencyKey: event.Header.IdempotencyKey,
	// }

	// resp, err := h.receiptsService.IssueReceipt(ctx, request)
	// if err != nil {
	// 	return fmt.Errorf("failed to issue receipt: %w", err)
	// }

	// return h.eventBus.Publish(ctx, entities.TicketReceiptIssued{
	// 	Header:        entities.NewEventHeaderWithIdempotencyKey(event.Header.IdempotencyKey),
	// 	TicketID:      event.TicketID,
	// 	ReceiptNumber: resp.ReceiptNumber,
	// 	IssuedAt:      resp.IssuedAt,
	// })
	return nil
}
