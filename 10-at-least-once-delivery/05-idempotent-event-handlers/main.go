package main

import "context"

type PaymentTaken struct {
	PaymentID string
	Amount    int
}

type PaymentsHandler struct {
	repo *PaymentsRepository
}

func NewPaymentsHandler(repo *PaymentsRepository) *PaymentsHandler {
	return &PaymentsHandler{repo: repo}
}

func (p *PaymentsHandler) HandlePaymentTaken(ctx context.Context, event *PaymentTaken) error {
	return p.repo.SavePaymentTaken(ctx, event)
}

type PaymentsRepository struct {
	payments       []PaymentTaken
	uniquePayments map[string]struct{}
}

func (p *PaymentsRepository) Payments() []PaymentTaken {
	return p.payments
}

func NewPaymentsRepository() *PaymentsRepository {
	return &PaymentsRepository{
		uniquePayments: make(map[string]struct{}),
	}
}

func (p *PaymentsRepository) SavePaymentTaken(ctx context.Context, event *PaymentTaken) error {
	if _, ok := p.uniquePayments[event.PaymentID]; !ok {
		p.payments = append(p.payments, *event)
		p.uniquePayments[event.PaymentID] = struct{}{}
	}
	return nil
}
