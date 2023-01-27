package model

import "time"

type TaskType string
type Frequency string

const (
	FieldCreatedAt = "createdAt"
)

const (
	CreateTransaction = TaskType("create-transaction")
)

const (
	Weekly  = Frequency("weekly")
	Monthly = Frequency("monthly")
	Yearly  = Frequency("yearly")
)

type RecurringTask struct {
	Id            string                 `json:"id"`
	UserID        string                 `json:"user"`
	Type          TaskType               `json:"type"`
	Frequency     Frequency              `json:"frequency"`
	RecurringTime uint                   `json:"recurringTime"`
	Body          map[string]interface{} `json:"body"`
	CreatedAt     time.Time              `json:"createdAt"`
}
