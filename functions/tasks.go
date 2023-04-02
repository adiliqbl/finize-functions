package functions

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/service"
	"finize-functions.app/util"
	"log"
	"sync"
)

var exchangeRates map[string]float64

func getExchangeRate(factory service.Factory, from string, to string) (float64, error) {
	if exchangeRates == nil {
		exchangeRates = map[string]float64{}
	}

	if rate, ok := exchangeRates[from+"_"+to]; ok {
		return rate, nil
	}

	if rate, err := GetExchangeRate(factory, from, to, false); err != nil {
		return 0, nil
	} else {
		exchangeRates[from+"_"+to] = rate.Rate
		return rate.Rate, nil
	}
}

func processRecurringTask(task model.RecurringTask, ops *[]data.DatabaseOperation, database service.FirestoreDB, factory service.Factory, clock util.Clock) bool {
	if task.Data == nil {
		log.Printf("No data found: %s", task.ID)
		return false
	}

	// Skip if already processed
	today, err := clock.WithZone(task.Timezone)
	if err != nil {
		return false
	}
	if task.LastDate != nil {
		if today.Year() == task.LastDate.UTC().Year() &&
			today.YearDay() == task.LastDate.UTC().YearDay() {
			return false
		}
	}

	// Check current date with the frequency
	switch task.Frequency {
	case model.Weekly:
		if int(today.Weekday()) != task.RecurringTime {
			return false
		}
		break
	case model.Monthly:
		if today.Day() != task.RecurringTime {
			return false
		}
		break
	case model.Yearly:
		if today.YearDay() != task.RecurringTime {
			return false
		}
		break
	}

	if task.Type == model.CreateTransaction {
		t := database.Collection(service.TransactionsDB(task.UserID)).NewDoc()

		date := today.UTC()
		task.Data[data.FieldId] = t.ID
		task.Data[model.FieldDate] = today.UTC()
		task.Data[model.FieldRecurringTask] = task.ID

		baseAmount, err := util.MapTo[model.Money](task.Data[model.FieldAmount])
		if err != nil {
			log.Printf("Failed to parse amount: %s: %v", task.ID, err)
			return false
		}

		amountLocal, err := util.MapTo[model.Money](task.Data[model.FieldAmountLocal])
		if err != nil {
			log.Printf("Failed to parse amountLocal: %s: %v", task.ID, err)
			return false
		}

		// Setting amountLocal
		if amountLocal.Currency != baseAmount.Currency {
			rate, err := getExchangeRate(factory, baseAmount.Currency, amountLocal.Currency)
			if err != nil {
				log.Printf("Failed to get exchange rate: %s: %v", task.ID, err)
				return false
			}
			amountLocal.Amount = rate * baseAmount.Amount
		} else {
			amountLocal.Amount = baseAmount.Amount
		}
		amountLocalMap, _ := util.MapTo[map[string]interface{}](amountLocal)
		task.Data[model.FieldAmountLocal] = amountLocalMap

		// Setting amountTo
		if accountTo, ok := task.Data[model.FieldAccountTo]; ok && accountTo != nil {
			amountTo, err := util.MapTo[model.Money](task.Data[model.FieldAmountTo])
			if err != nil {
				log.Printf("Failed to parse amountTo: %s: %v", task.ID, err)
				return false
			}

			if amountTo.Currency != baseAmount.Currency {
				rate, err := getExchangeRate(factory, baseAmount.Currency, amountTo.Currency)
				if err != nil {
					log.Printf("Failed to get exchange rate: %s: %v", task.ID, err)
					return false
				}
				amountTo.Amount = rate * baseAmount.Amount
			} else {
				amountTo.Amount = baseAmount.Amount
			}

			amountToMap, _ := util.MapTo[map[string]interface{}](amountTo)
			task.Data[model.FieldAmountTo] = amountToMap
		} else {
			task.Data[model.FieldAmountTo] = nil
		}

		// Setting amountFrom
		if accountFrom, ok := task.Data[model.FieldAccountFrom]; ok && accountFrom != nil {
			amountFrom, err := util.MapTo[model.Money](task.Data[model.FieldAmountFrom])
			if err != nil {
				log.Printf("Failed to parse amountTo: %s: %v", task.ID, err)
				return false
			}

			if amountFrom.Currency != baseAmount.Currency {
				rate, err := getExchangeRate(factory, baseAmount.Currency, amountFrom.Currency)
				if err != nil {
					log.Printf("Failed to get exchange rate: %s: %v", task.ID, err)
					return false
				}
				amountFrom.Amount = rate * baseAmount.Amount
			} else {
				amountFrom.Amount = baseAmount.Amount
			}

			amountFromMap, _ := util.MapTo[map[string]interface{}](amountFrom)
			task.Data[model.FieldAmountFrom] = amountFromMap
		} else {
			task.Data[model.FieldAmountFrom] = nil
		}

		*ops = append(*ops, data.DatabaseOperation{
			Ref:  database.Doc(service.TransactionDoc(task.UserID, t.ID)),
			Data: task.Data,
		})

		*ops = append(*ops, data.DatabaseOperation{
			Ref: database.Doc(service.TaskDoc(task.ID)),
			Data: firestore.Update{
				Path:  model.FieldLastDate,
				Value: date,
			},
		})
	}

	return true
}

func ProcessRecurringTasks(factory service.Factory, clock util.Clock) error {
	database := factory.Firestore()
	tasks := factory.TaskService()

	const batchSize = 250
	offset := 0

	for true {
		list, err := tasks.Paginate(offset, batchSize)
		if err != nil {
			log.Printf("Failed to get tasks with offset %d: %v", offset, err)
			return err
		}

		var ops []data.DatabaseOperation

		var wg = &sync.WaitGroup{}
		wg.Add(len(list))
		for _, task := range list {
			go func(task model.RecurringTask) {
				defer wg.Done()
				processRecurringTask(task, &ops, database, factory, clock)
			}(task)
		}
		wg.Wait()

		if len(ops) > 0 {
			err = database.Batch(func() []data.DatabaseOperation {
				return ops
			})

			if err != nil {
				log.Printf("Failed to write batch: %v", err)
				return err
			}
		}

		if len(list) < batchSize {
			break
		}

		offset += batchSize
	}

	return nil
}
