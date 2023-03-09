package functions

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/service"
	"log"
)

func OnUserUpdated(factory service.Factory, oldUser model.User, user model.User, _ []string) error {
	if oldUser.Profile.Currency == user.Profile.Currency {
		return nil
	}

	tasks := factory.TaskService()

	return factory.Firestore().Transaction(func(tx *firestore.Transaction) []data.DatabaseOperation {
		var ops []data.DatabaseOperation

		userTasks, err := tasks.FindByUser(user.ID)
		if err != nil {
			log.Fatalf("AccountService.FindByBudget: %v", err)
		}

		for _, task := range userTasks {
			ops = append(ops, data.DatabaseOperation{
				Ref: tasks.Doc(task.ID),
				Data: []firestore.Update{{
					Path:  model.FieldData + "." + model.FieldAmountLocal + "." + model.FieldCurrency,
					Value: user.Profile.Currency,
				}},
			})
		}

		return ops
	})
}
