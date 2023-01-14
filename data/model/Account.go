package model

import . "finize-functions.app/data"

type AccountEvent struct {
	ID       StringValue    `json:"id"`
	Name     StringValue    `json:"name"`
	Balance  DoubleValue    `json:"balance"`
	Type     StringValue    `json:"type"`
	Currency StringValue    `json:"currency"`
	Budget   ReferenceValue `json:"budget,omitempty"`
}

type Account struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Balance  float64 `json:"balance"`
	Type     string  `json:"type"`
	Currency string  `json:"currency"`
	Budget   *string `json:"budget,omitempty"`
}
