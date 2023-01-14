package model

import (
	. "finize-functions.app/data"
	"time"
)

type TransactionEvent struct {
	ID          StringValue        `json:"id"`
	Name        StringValue        `json:"name"`
	Amount      DoubleValue        `json:"amount"`
	AmountValue DoubleValue        `json:"amountValue,omitempty"`
	AccountTo   ReferenceValue     `json:"accountTo,omitempty"`
	AccountFrom ReferenceValue     `json:"accountFrom,omitempty"`
	Budget      ReferenceValue     `json:"budget,omitempty"`
	Currency    StringValue        `json:"currency"`
	Category    ArrayValue[string] `json:"category"`
	Date        TimestampValue     `json:"date"`
}

type Transaction struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Amount      float64   `json:"amount"`
	AmountValue *float64  `json:"amountValue"`
	AccountTo   *string   `json:"accountTo,omitempty"`
	AccountFrom *string   `json:"accountFrom,omitempty"`
	Budget      *string   `json:"budget,omitempty"`
	Currency    string    `json:"currency"`
	Category    []string  `json:"category"`
	Date        time.Time `json:"date"`
}
