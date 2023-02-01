package model

import "time"

type ExchangeRate struct {
	Rate float64   `json:"rate"`
	Date time.Time `json:"date"`
}
