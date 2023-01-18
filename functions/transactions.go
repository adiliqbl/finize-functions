package functions

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data/model"
	"finize-functions.app/service"
	"finize-functions.app/util"
	"log"
)

func OnTransactionCreated(userID string, transaction model.Transaction) error {
	store := service.NewFirestoreService()
	accounts := service.NewAccountService(userID)
	budgets := service.NewBudgetService(userID)

	return store.Transaction(func(tx *firestore.Transaction) error {
		if !util.NullOrEmpty(transaction.AccountTo) {
			budget, err := budgets.FindByIDWith(*transaction.Budget, tx)
			if err != nil {
				log.Fatalf("BudgetService.FindByID: %v", err)
			}
			budget.Spent = budget.Spent + transaction.AmountLocal.Amount

			doc := budgets.Doc(budget.ID)
			err = tx.Update(doc, []firestore.Update{{
				Path:  model.FieldSpent,
				Value: budget.Spent,
			}})

			if err != nil {
				log.Fatalf("BudgetTransaction.Update: %v", err)
				return err
			}
		}

		if !util.NullOrEmpty(transaction.AccountFrom) {
			accountFrom, err := accounts.FindByIDWith(*transaction.AccountFrom, tx)
			if err != nil {
				log.Fatalf("AccountService.FindByID: %v", err)
			}
			accountFrom.Balance = accountFrom.Balance - transaction.AmountFrom.Amount

			doc := accounts.Doc(accountFrom.ID)
			err = tx.Update(doc, []firestore.Update{{
				Path:  model.FieldBalance,
				Value: accountFrom.Balance,
			}})

			if err != nil {
				log.Fatalf("AccountTransaction.Update: %v", err)
				return err
			}
		}

		if !util.NullOrEmpty(transaction.Budget) {
			accountTo, err := accounts.FindByIDWith(*transaction.AccountTo, tx)
			if err != nil {
				log.Fatalf("AccountService.FindByID: %v", err)
			}
			accountTo.Balance = accountTo.Balance - transaction.AmountTo.Amount

			doc := accounts.Doc(accountTo.ID)
			err = tx.Update(doc, []firestore.Update{{
				Path:  model.FieldBalance,
				Value: accountTo.Balance,
			}})

			if err != nil {
				log.Fatalf("AccountTransaction.Update: %v", err)
				return err
			}
		}

		return nil
	})
}
