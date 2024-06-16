package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type IssueReceiptRequest struct {
	TicketID string `json:"ticket_id"`
	Price    Money  `json:"price"`
}

type Money struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type IssueReceiptResponse struct {
	ReceiptNumber string    `json:"number"`
	IssuedAt      time.Time `json:"issued_at"`
}

type ReceiptsService interface {
	IssueReceipt(ctx context.Context, request IssueReceiptRequest) (IssueReceiptResponse, error)
}

type ReceiptsServiceMock struct {
	mock           sync.Mutex
	lastReceiptNum int
	IssuedReceipts []IssueReceiptRequest
}

func (rsm *ReceiptsServiceMock) IssueReceipt(ctx context.Context, req IssueReceiptRequest) (IssueReceiptResponse, error) {
	rsm.mock.Lock()
	defer rsm.mock.Unlock()
	rsm.lastReceiptNum++
	rsm.IssuedReceipts = append(rsm.IssuedReceipts, req)
	return IssueReceiptResponse{
		ReceiptNumber: fmt.Sprint(rsm.lastReceiptNum),
		IssuedAt:      time.Now(),
	}, nil
}
