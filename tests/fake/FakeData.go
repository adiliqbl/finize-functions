package fake

import (
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"time"
)

func MapTo[T any](object any) T {
	obj, _ := util.MapTo[T](object)
	return obj
}

func NewUser(obj model.UserEvent) model.User {
	return MapTo[model.User](obj)
}

func NewUserEvent(id string, name string, email string) model.UserEvent {
	return model.UserEvent{
		ID:    data.StringValue{Value: &id},
		Name:  data.StringValue{Value: &name},
		Email: data.StringValue{Value: &email},
	}
}

func NewUserEventMap(user model.UserEvent) map[string]interface{} {
	return map[string]interface{}{
		"id": map[string]interface{}{
			"stringValue": user.ID.Value,
		},
		"name": map[string]interface{}{
			"stringValue": user.Name.Value,
		},
		"email": map[string]interface{}{
			"stringValue": user.Email.Value,
		},
	}
}

func NewBudget(obj model.BudgetEvent) model.Budget {
	return MapTo[model.Budget](obj)
}

func NewBudgetEvent(id string, name string, limit float64) model.BudgetEvent {
	return model.BudgetEvent{
		ID:    data.StringValue{Value: &id},
		Name:  data.StringValue{Value: &name},
		Limit: data.DoubleValue{Value: &limit},
	}
}

func NewBudgetEventMap(budget model.BudgetEvent) map[string]interface{} {
	return map[string]interface{}{
		"id": map[string]interface{}{
			"stringValue": budget.ID.Value,
		},
		"name": map[string]interface{}{
			"stringValue": budget.Name.Value,
		},
		"limit": map[string]interface{}{
			"doubleValue": budget.Limit.Value,
		},
	}
}

func NewAccount(obj model.AccountEvent) model.Account {
	return MapTo[model.Account](obj)
}

func NewAccountEvent(id string, name string, balance float64, budget *string) model.AccountEvent {
	return model.AccountEvent{
		ID:       data.StringValue{Value: &id},
		Name:     data.StringValue{Value: &name},
		Balance:  data.DoubleValue{Value: &balance},
		Currency: data.StringValue{Value: util.Pointer("CURR")},
		Type:     data.StringValue{Value: util.Pointer("type")},
		Budget:   data.ReferenceValue{Reference: budget},
	}
}

func NewAccountEventMap(account model.AccountEvent) map[string]interface{} {
	return map[string]interface{}{
		"id": map[string]interface{}{
			"stringValue": account.ID.Value,
		},
		"name": map[string]interface{}{
			"stringValue": account.Name.Value,
		},
		"type": map[string]interface{}{
			"stringValue": account.Type.Value,
		},
		"balance": map[string]interface{}{
			"doubleValue": account.Balance.Value,
		},
		"currency": map[string]interface{}{
			"stringValue": account.Currency.Value,
		},
		"budget": map[string]interface{}{
			"stringValue": util.ValueOrNull(account.Budget.Get()),
		},
	}
}

func NewTransaction(obj model.TransactionEvent) model.Transaction {
	return MapTo[model.Transaction](obj)
}

func NewTransactionEvent(id string, name string, amount float64, amountLocal float64, amountValue float64, date time.Time, accountTo *string, accountFrom *string, budget *string) model.TransactionEvent {
	transaction := model.TransactionEvent{
		ID:       data.StringValue{Value: &id},
		Name:     data.StringValue{Value: &name},
		Category: data.ArrayValue[string]{Value: &[]string{"One", "Two"}},
		Date:     data.TimestampValue{Value: &date},
		Amount: data.MapValue[model.MoneyEvent]{Value: util.Pointer(model.MoneyEvent{
			Amount:   data.DoubleValue{Value: &amount},
			Currency: data.StringValue{Value: util.Pointer("CURR")},
		})},
		AmountLocal: data.MapValue[model.MoneyEvent]{Value: util.Pointer(model.MoneyEvent{
			Amount:   data.DoubleValue{Value: &amountLocal},
			Currency: data.StringValue{Value: util.Pointer("CURR")},
		})},
		AmountTo: data.MapValue[model.MoneyEvent]{Value: util.Pointer(model.MoneyEvent{
			Amount:   data.DoubleValue{Value: &amountValue},
			Currency: data.StringValue{Value: util.Pointer("CURR")},
		})},
		AmountFrom: data.MapValue[model.MoneyEvent]{Value: util.Pointer(model.MoneyEvent{
			Amount:   data.DoubleValue{Value: &amountValue},
			Currency: data.StringValue{Value: util.Pointer("CURR")},
		})},
		AccountTo:   data.ReferenceValue{Reference: accountTo},
		AccountFrom: data.ReferenceValue{Reference: accountFrom},
		Budget:      data.ReferenceValue{Reference: budget},
	}

	if accountFrom != nil {
		transaction.AmountFrom = data.MapValue[model.MoneyEvent]{Value: util.Pointer(model.MoneyEvent{
			Amount:   data.DoubleValue{Value: &amountValue},
			Currency: data.StringValue{Value: util.Pointer("CURR")},
		})}
	}

	if accountTo != nil {
		transaction.AmountTo = data.MapValue[model.MoneyEvent]{Value: util.Pointer(model.MoneyEvent{
			Amount:   data.DoubleValue{Value: &amountValue},
			Currency: data.StringValue{Value: util.Pointer("CURR")},
		})}
	}

	return transaction
}

func NewTransactionEventMap(transaction model.TransactionEvent) map[string]interface{} {
	doc := map[string]interface{}{
		"id": map[string]interface{}{
			"stringValue": transaction.ID.Value,
		},
		"name": map[string]interface{}{
			"stringValue": transaction.Name.Value,
		},
		"category": map[string]interface{}{
			"arrayValue": transaction.Category.Value,
		},
		"date": map[string]interface{}{
			"timestampValue": transaction.Date.Value,
		},
		"amount": map[string]interface{}{
			"mapValue": map[string]interface{}{
				"amount": map[string]interface{}{
					"doubleValue": transaction.Amount.Value.Amount,
				},
				"currency": map[string]interface{}{
					"stringValue": transaction.Amount.Value.Currency,
				},
			},
		},
		"amountLocal": map[string]interface{}{
			"mapValue": map[string]interface{}{
				"amount": map[string]interface{}{
					"doubleValue": transaction.AmountLocal.Value.Amount,
				},
				"currency": map[string]interface{}{
					"stringValue": transaction.AmountLocal.Value.Currency,
				},
			},
		},
		"amountFrom": map[string]interface{}{
			"mapValue": map[string]interface{}{
				"amount": map[string]interface{}{
					"doubleValue": transaction.AmountFrom.Value.Amount,
				},
				"currency": map[string]interface{}{
					"stringValue": transaction.AmountFrom.Value.Currency,
				},
			},
		},
		"amountTo": map[string]interface{}{
			"mapValue": map[string]interface{}{
				"amount": map[string]interface{}{
					"doubleValue": transaction.AmountTo.Value.Amount,
				},
				"currency": map[string]interface{}{
					"stringValue": transaction.AmountTo.Value.Currency,
				},
			},
		},
		"accountTo": map[string]interface{}{
			"stringValue": transaction.AccountTo.Get(),
		},
		"accountFrom": map[string]interface{}{
			"stringValue": transaction.AccountFrom.Get(),
		},
		"budget": map[string]interface{}{
			"stringValue": transaction.Budget.Get(),
		},
	}

	if transaction.AmountFrom.Value != nil {
		doc["amountFrom"] = map[string]interface{}{
			"mapValue": map[string]interface{}{
				"amount": map[string]interface{}{
					"doubleValue": transaction.AmountFrom.Value.Amount,
				},
				"currency": map[string]interface{}{
					"stringValue": transaction.AmountFrom.Value.Currency,
				},
			},
		}
	} else {
		doc["amountFrom"] = map[string]interface{}{
			"mapValue": nil,
		}
	}

	if transaction.AmountTo.Value != nil {
		doc["amountTo"] = map[string]interface{}{
			"mapValue": map[string]interface{}{
				"amount": map[string]interface{}{
					"doubleValue": transaction.AmountTo.Value.Amount,
				},
				"currency": map[string]interface{}{
					"stringValue": transaction.AmountTo.Value.Currency,
				},
			},
		}
	} else {
		doc["amountTo"] = map[string]interface{}{
			"mapValue": nil,
		}
	}

	return doc
}

func NewRecurringTask(userID string, task model.TaskType, recurringTime int, frequency model.Frequency, timezone string, lastDate *time.Time, body map[string]interface{}) model.RecurringTask {
	return model.RecurringTask{
		UserID:        userID,
		Type:          task,
		RecurringTime: recurringTime,
		Timezone:      timezone,
		Frequency:     frequency,
		Data:          body,
		LastDate:      lastDate,
		CreatedAt:     time.Now().UTC(),
	}
}
