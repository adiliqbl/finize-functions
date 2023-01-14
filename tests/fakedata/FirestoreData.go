package fakedata

import (
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"time"
)

func NewUserEvent(id string, name string, email string) model.UserEvent {
	return model.UserEvent{
		ID:    data.StringValue{Value: util.Pointer(id)},
		Name:  data.StringValue{Value: util.Pointer(name)},
		Email: data.StringValue{Value: util.Pointer(email)},
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

func NewBudgetEvent(id string, name string, limit float64, spent float64) model.BudgetEvent {
	return model.BudgetEvent{
		ID:    data.StringValue{Value: util.Pointer(id)},
		Name:  data.StringValue{Value: util.Pointer(name)},
		Limit: data.DoubleValue{Value: util.Pointer(limit)},
		Spent: data.DoubleValue{Value: &spent},
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
		"spent": map[string]interface{}{
			"doubleValue": budget.Spent.Value,
		},
	}
}

func NewAccountEvent(id string, name string, balance float64, budget *string) model.AccountEvent {
	return model.AccountEvent{
		ID:       data.StringValue{Value: util.Pointer(id)},
		Name:     data.StringValue{Value: util.Pointer(name)},
		Balance:  data.DoubleValue{Value: util.Pointer(balance)},
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
			"referenceValue": util.ValueOrNull(account.Budget.Get()),
		},
	}
}

func NewTransactionEvent(id string, name string, amount float64, amountValue *float64, date time.Time, accountTo *string, accountFrom *string, budget *string) model.TransactionEvent {
	return model.TransactionEvent{
		ID:          data.StringValue{Value: util.Pointer(id)},
		Name:        data.StringValue{Value: util.Pointer(name)},
		Amount:      data.DoubleValue{Value: util.Pointer(amount)},
		Currency:    data.StringValue{Value: util.Pointer("CURR")},
		Category:    data.ArrayValue[string]{Value: &[]string{"One", "Two"}},
		Date:        data.TimestampValue{Value: util.Pointer(time.Now())},
		AmountValue: data.DoubleValue{Value: amountValue},
		AccountTo:   data.ReferenceValue{Reference: accountTo},
		AccountFrom: data.ReferenceValue{Reference: accountFrom},
		Budget:      data.ReferenceValue{Reference: budget},
	}
}

func NewTransactionEventMap(transaction model.TransactionEvent) map[string]interface{} {
	return map[string]interface{}{
		"id": map[string]interface{}{
			"stringValue": transaction.ID.Value,
		},
		"name": map[string]interface{}{
			"stringValue": transaction.Name.Value,
		},
		"amount": map[string]interface{}{
			"doubleValue": transaction.Amount.Value,
		},
		"currency": map[string]interface{}{
			"stringValue": "CURR",
		},
		"category": map[string]interface{}{
			"arrayValue": transaction.Category.Value,
		},
		"date": map[string]interface{}{
			"timestampValue": transaction.Date.Value,
		},
		"amountValue": map[string]interface{}{
			"doubleValue": transaction.AmountValue,
		},
		"accountTo": map[string]interface{}{
			"referenceValue": transaction.AccountTo.Get(),
		},
		"accountFrom": map[string]interface{}{
			"referenceValue": transaction.AccountFrom.Get(),
		},
		"budget": map[string]interface{}{
			"referenceValue": transaction.Budget.Get(),
		},
	}
}
