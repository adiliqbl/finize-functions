package service

import (
	"context"
	"finize-functions.app/data/model"
)

type Factory interface {
	Firestore() FirestoreDB
	EventService() EventService
	UserService() UserService
	AccountService() AccountService
	BudgetService() BudgetService
	TransactionService() TransactionService
	ExchangeRateService() ExchangeRateService
}

type serviceProvider struct {
	ctx     context.Context
	userID  string
	eventID string
}

func NewServiceFactory(ctx context.Context, eventID string, userID string) Factory {
	return &serviceProvider{ctx: ctx, eventID: eventID, userID: userID}
}

func (p *serviceProvider) Firestore() FirestoreDB {
	return newFirestoreDB(p.ctx, p.eventID)
}

func (p *serviceProvider) EventService() EventService {
	return NewEventService(newFirestoreService[model.Event](p.ctx, p.eventID), p.eventID)
}

func (p *serviceProvider) UserService() UserService {
	return NewUserService(newFirestoreService[model.User](p.ctx, p.eventID))
}

func (p *serviceProvider) AccountService() AccountService {
	return NewAccountService(newFirestoreService[model.Account](p.ctx, p.eventID), p.userID)
}

func (p *serviceProvider) BudgetService() BudgetService {
	return NewBudgetService(newFirestoreService[model.Budget](p.ctx, p.eventID), p.userID)
}

func (p *serviceProvider) TransactionService() TransactionService {
	return NewTransactionService(newFirestoreService[model.Transaction](p.ctx, p.eventID), p.userID)
}

func (p *serviceProvider) ExchangeRateService() ExchangeRateService {
	return NewExchangeRateService(newFirestoreService[map[string]interface{}](p.ctx, p.eventID))
}
