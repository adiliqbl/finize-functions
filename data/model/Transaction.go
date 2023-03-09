package model

import (
	. "finize-functions.app/data"
	"time"
)

const (
	FieldName          = "name"
	FieldDate          = "date"
	FieldRecurringTask = "recurringTask"
	FieldAccountTo     = "accountTo"
	FieldAccountFrom   = "accountFrom"
	FieldAmount        = "amount"
	FieldAmountTo      = "amountTo"
	FieldAmountFrom    = "amountFrom"
	FieldAmountLocal   = "amountLocal"
)

type TransactionEvent struct {
	ID          StringValue          `json:"id"`
	Name        StringValue          `json:"name"`
	Amount      MapValue[MoneyEvent] `json:"amount"`
	AmountTo    MapValue[MoneyEvent] `json:"amountTo,omitempty"`
	AmountFrom  MapValue[MoneyEvent] `json:"amountFrom,omitempty"`
	AmountLocal MapValue[MoneyEvent] `json:"amountLocal,omitempty"`
	AccountTo   ReferenceValue       `json:"accountTo,omitempty"`
	AccountFrom ReferenceValue       `json:"accountFrom,omitempty"`
	Budget      ReferenceValue       `json:"budget,omitempty"`
	Task        ReferenceValue       `json:"recurringTask,omitempty"`
	Categories  ArrayValue[string]   `json:"categories"`
	Date        TimestampValue       `json:"date"`
}

type Transaction struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Amount      Money     `json:"amount"`
	AmountTo    *Money    `json:"amountTo,omitempty"`
	AmountFrom  *Money    `json:"amountFrom,omitempty"`
	AmountLocal Money     `json:"amountLocal,omitempty"`
	AccountTo   *string   `json:"accountTo,omitempty"`
	AccountFrom *string   `json:"accountFrom,omitempty"`
	Budget      *string   `json:"budget,omitempty"`
	Task        *string   `json:"recurringTask,omitempty"`
	Categories  []string  `json:"categories"`
	Date        time.Time `json:"date"`
}
