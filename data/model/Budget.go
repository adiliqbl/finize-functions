package model

import . "finize-functions.app/data"

type BudgetEvent struct {
	ID    StringValue `json:"id"`
	Name  StringValue `json:"name"`
	Limit DoubleValue `json:"limit"`
	Spent DoubleValue `json:"spent,omitempty"`
}

type Budget struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Limit float64  `json:"limit"`
	Spent *float64 `json:"spent,omitempty"`
}
