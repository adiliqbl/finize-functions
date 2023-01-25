package model

import "finize-functions.app/data"

type RecurringTask struct {
	UserID    string                 `json:"user"`
	Type      data.TaskType          `json:"type"`
	Frequency int                    `json:"frequency"`
	Body      map[string]interface{} `json:"body"`
}
