package fake

import (
	"context"
	"finize-functions.app/data/model"
	"finize-functions.app/service"
)

type provider struct {
	ctx    context.Context
	userID string
}

func NewServiceFactory(ctx context.Context, userID string) service.Factory {
	return &provider{ctx: ctx, userID: userID}
}

func (p *provider) FirestoreService() service.FirestoreService[model.Entity] {
	return NewFirestoreService[model.Entity](p.ctx)
}

func (p *provider) UserService() service.UserService {
	return service.NewUserService(NewFirestoreService[model.User](p.ctx))
}

func (p *provider) AccountService() service.AccountService {
	return service.NewAccountService(NewFirestoreService[model.Account](p.ctx), p.userID)
}

func (p *provider) BudgetService() service.BudgetService {
	return service.NewBudgetService(NewFirestoreService[model.Budget](p.ctx), p.userID)
}

func (p *provider) TransactionService() service.TransactionService {
	return service.NewTransactionService(NewFirestoreService[model.Transaction](p.ctx), p.userID)
}
