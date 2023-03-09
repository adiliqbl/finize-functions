package functions

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/service"
	"log"
)

func OnAccountUpdated(factory service.Factory, oldAccount model.Account, account model.Account, _ []string) error {
	if oldAccount.Currency == account.Currency {
		return nil
	}

	tasks := factory.TaskService()

	return factory.Firestore().Transaction(func(tx *firestore.Transaction) []data.DatabaseOperation {
		var ops []data.DatabaseOperation

		accountTasks, err := tasks.FindByAccount(account.ID)
		if err != nil {
			log.Fatalf("AccountService.FindByBudget: %v", err)
		}

		for _, task := range accountTasks {
			if _, ok := task.Data[model.FieldAccountTo]; ok {
				ops = append(ops, data.DatabaseOperation{
					Ref: tasks.Doc(task.ID),
					Data: []firestore.Update{{
						Path:  model.FieldData + "." + model.FieldAmountTo + "." + model.FieldCurrency,
						Value: account.Currency,
					}},
				})
			}

			if _, ok := task.Data[model.FieldAccountFrom]; ok {
				ops = append(ops, data.DatabaseOperation{
					Ref: tasks.Doc(task.ID),
					Data: []firestore.Update{{
						Path:  model.FieldData + "." + model.FieldAmountFrom + "." + model.FieldCurrency,
						Value: account.Currency,
					}},
				})
			}
		}

		return ops
	})
}
