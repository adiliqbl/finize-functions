package functions

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/service"
	"finize-functions.app/util"
	"log"
)

func OnTransactionCreated(factory service.Factory, transaction model.Transaction) error {
	store := factory.FirestoreService()
	accounts := factory.AccountService()
	budgets := factory.BudgetService()

	return store.Transaction(func(tx *firestore.Transaction) error {
		var ops []data.TransactionOperation

		if !util.NullOrEmpty(transaction.AccountTo) {
			budget, err := budgets.FindByIDWith(*transaction.Budget, tx)
			if err != nil {
				log.Fatalf("BudgetService.FindByID: %v", err)
			}
			budget.Spent = budget.Spent + transaction.AmountLocal.Amount

			ops = append(ops, data.TransactionOperation{
				Ref: budgets.Doc(budget.ID),
				Data: []firestore.Update{{
					Path:  model.FieldSpent,
					Value: budget.Spent,
				}},
			})
		}

		if !util.NullOrEmpty(transaction.AccountFrom) {
			accountFrom, err := accounts.FindByIDWith(*transaction.AccountFrom, tx)
			if err != nil {
				log.Fatalf("AccountService.FindByID: %v", err)
			}
			accountFrom.Balance = accountFrom.Balance - transaction.AmountFrom.Amount

			ops = append(ops, data.TransactionOperation{
				Ref: accounts.Doc(accountFrom.ID),
				Data: []firestore.Update{{
					Path:  model.FieldBalance,
					Value: accountFrom.Balance,
				}},
			})
		}

		if !util.NullOrEmpty(transaction.Budget) {
			accountTo, err := accounts.FindByIDWith(*transaction.AccountTo, tx)
			if err != nil {
				log.Fatalf("AccountService.FindByID: %v", err)
			}
			accountTo.Balance = accountTo.Balance - transaction.AmountTo.Amount

			ops = append(ops, data.TransactionOperation{
				Ref: accounts.Doc(accountTo.ID),
				Data: []firestore.Update{{
					Path:  model.FieldBalance,
					Value: accountTo.Balance,
				}},
			})
		}

		return data.Commit(tx, ops)
	})
}
