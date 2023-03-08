package functions

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/service"
	"log"
)

func OnBudgetDeleted(factory service.Factory, budget model.Budget) error {
	accounts := factory.AccountService()

	return factory.Firestore().Transaction(func(tx *firestore.Transaction) []data.DatabaseOperation {
		var ops []data.DatabaseOperation

		budgetAccounts, err := accounts.FindByBudget(budget.ID)
		if err != nil {
			log.Fatalf("AccountService.FindByBudget: %v", err)
		}

		for _, account := range budgetAccounts {
			ops = append(ops, data.DatabaseOperation{
				Ref: accounts.Doc(account.ID),
				Data: []firestore.Update{{
					Path:  model.FieldBudget,
					Value: nil,
				}},
			})
		}

		return ops
	})
}
