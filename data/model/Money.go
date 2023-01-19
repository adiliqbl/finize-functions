package model

import (
	. "finize-functions.app/data"
)

type MoneyEvent struct {
	Amount   DoubleValue `json:"amount"`
	Currency StringValue `json:"currency"`
}

type Money struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
