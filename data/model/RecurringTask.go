package model

import "time"

type TaskType string

const (
	FieldCreatedAt = "createdAt"
)

const (
	CreateTransaction = TaskType("create-transaction")
)

type RecurringTask struct {
	Id        string                 `json:"id"`
	UserID    string                 `json:"user"`
	Type      TaskType               `json:"type"`
	Frequency uint                   `json:"frequency"`
	Body      map[string]interface{} `json:"body"`
	CreatedAt time.Time              `json:"createdAt"`
}
