package functions

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/service"
	"finize-functions.app/util"
	"log"
)

func ProcessRecurringTasks(factory service.Factory, clock util.Clock) error {
	database := factory.Firestore()
	tasks := factory.TaskService()

	offset := 0

	for true {
		list, err := tasks.Paginate(offset, 250)
		if err != nil {
			log.Printf("Failed to get tasks with offset %d: %v", offset, err)
			return err
		}

		err = database.Batch(func() []data.DatabaseOperation {
			var ops []data.DatabaseOperation

			for _, task := range list {
				if task.Data == nil {
					log.Printf("No data found: %v", err)
					continue
				}

				// Skip if already processed
				today, err := clock.WithZone(task.Timezone)
				if err != nil {
					continue
				}
				if task.LastDate != nil {
					if today.Year() == task.LastDate.UTC().Year() &&
						today.YearDay() == task.LastDate.UTC().YearDay() {
						continue
					}
				}

				// Check current date with the frequency
				switch task.Frequency {
				case model.Weekly:
					if int(today.Weekday()) != task.RecurringTime {
						continue
					}
					break
				case model.Monthly:
					if today.Day() != task.RecurringTime {
						continue
					}
					break
				case model.Yearly:
					if today.YearDay() != task.RecurringTime {
						continue
					}
					break
				}

				if task.Type == model.CreateTransaction {
					t := database.Collection(service.TransactionsDB(task.UserID)).NewDoc()

					date := today.UTC()
					task.Data[data.FieldId] = t.ID
					task.Data[model.FieldDate] = today.UTC()
					task.Data[model.FieldRecurringTask] = task.Id

					ops = append(ops, data.DatabaseOperation{
						Ref:  database.Doc(service.TransactionDoc(task.UserID, t.ID)),
						Data: task.Data,
					})

					ops = append(ops, data.DatabaseOperation{
						Ref: database.Doc(service.TaskDoc(task.Id)),
						Data: firestore.Update{
							Path:  model.FieldLastDate,
							Value: date,
						},
					})
				}
			}

			return ops
		})

		if err != nil {
			log.Printf("Failed to write batch: %v", err)
			return err
		}

		if len(list) < 250 {
			break
		} else {
			offset = offset + 250
		}
	}

	return nil
}
