package fake

import (
	"context"
	"finize-functions.app/data/model"
	"finize-functions.app/service"
)

type serviceProvider struct {
	ctx    context.Context
	userID string
}

func NewServiceFactory(ctx context.Context, userID string) service.Factory {
	return &serviceProvider{ctx: ctx, userID: userID}
}

func (p *serviceProvider) Firestore() service.FirestoreDB {
	return NewFirestoreDB(p.ctx)
}

func (p *serviceProvider) EventService() service.EventService {
	return service.NewEventService(NewFirestoreService[model.Event](p.ctx), "event")
}

func (p *serviceProvider) UserService() service.UserService {
	return service.NewUserService(NewFirestoreService[model.User](p.ctx))
}

func (p *serviceProvider) AccountService() service.AccountService {
	return service.NewAccountService(NewFirestoreService[model.Account](p.ctx), p.userID)
}

func (p *serviceProvider) BudgetService() service.BudgetService {
	return service.NewBudgetService(NewFirestoreService[model.Budget](p.ctx), p.userID)
}

func (p *serviceProvider) TransactionService() service.TransactionService {
	return service.NewTransactionService(NewFirestoreService[model.Transaction](p.ctx), p.userID)
}
