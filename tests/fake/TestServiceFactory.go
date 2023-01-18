package fake

import (
	"finize-functions.app/data/model"
	"finize-functions.app/service"
	"github.com/stretchr/testify/mock"
)

type TestServiceFactory struct {
	mock.Mock
}

func (p *TestServiceFactory) FirestoreService() service.FirestoreService[model.Entity] {
	return &TestFirestore[model.Entity]{}
}

func (p *TestServiceFactory) UserService() service.UserService {
	return &TestService[model.User]{}
}

func (p *TestServiceFactory) AccountService() service.AccountService {
	return &TestService[model.Account]{}
}

func (p *TestServiceFactory) BudgetService() service.BudgetService {
	return &TestService[model.Budget]{}
}

func (p *TestServiceFactory) TransactionService() service.TransactionService {
	return &TestService[model.Transaction]{}
}
