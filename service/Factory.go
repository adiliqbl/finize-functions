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
	return &userServiceImpl{db: newFirestoreService[model.User](p.ctx)}
}

func (p *provider) AccountService() AccountService {
	return &accountServiceImpl{db: newFirestoreService[model.Account](p.ctx), userID: p.userID}
}

func (p *provider) BudgetService() BudgetService {
	return &budgetServiceImpl{db: newFirestoreService[model.Budget](p.ctx), userID: p.userID}
}

func (p *provider) TransactionService() TransactionService {
	return &transactionServiceImpl{db: newFirestoreService[model.Transaction](p.ctx), userID: p.userID}
}
