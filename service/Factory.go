package service

import (
	"context"
	"finize-functions.app/data/model"
)

type Factory interface {
	FirestoreService() FirestoreService[model.Entity]
	UserService() UserService
	AccountService() AccountService
	BudgetService() BudgetService
	TransactionService() TransactionService
}

type provider struct {
	ctx    context.Context
	userID string
}

func NewServiceFactory(ctx context.Context, userID string) Factory {
	return &provider{ctx: ctx, userID: userID}
}

func (p *provider) FirestoreService() FirestoreService[model.Entity] {
	return newFirestoreService[model.Entity](p.ctx)
}

func (p *provider) UserService() UserService {
	return NewUserService(newFirestoreService[model.User](p.ctx))
}

func (p *provider) AccountService() AccountService {
	return NewAccountService(newFirestoreService[model.Account](p.ctx), p.userID)
}

func (p *provider) BudgetService() BudgetService {
	return NewBudgetService(newFirestoreService[model.Budget](p.ctx), p.userID)
}

func (p *provider) TransactionService() TransactionService {
	return NewTransactionService(newFirestoreService[model.Transaction](p.ctx), p.userID)
}
