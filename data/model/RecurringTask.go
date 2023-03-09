package model

import "time"

type TaskType string
type Frequency string

const (
	FieldUser      = "user"
	FieldData      = "data"
	FieldLastDate  = "lastDate"
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
	ID            string                 `json:"id"`
	UserID        string                 `json:"user"`
	Type          TaskType               `json:"type"`
	Frequency     Frequency              `json:"frequency"`
	RecurringTime int                    `json:"recurringTime"`
	Timezone      string                 `json:"timezone"`
	Data          map[string]interface{} `json:"data"`
	LastDate      *time.Time             `json:"lastDate,omitempty"`
	CreatedAt     time.Time              `json:"createdAt"`
}
