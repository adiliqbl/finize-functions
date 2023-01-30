package service

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data/model"
	"fmt"
)

type TaskService interface {
	Paginate(start int, limit int) ([]model.RecurringTask, error)
}

type taskServiceImpl struct {
	db FirestoreService[model.RecurringTask]
}

func TasksDB() string {
	return "tasks"
}

func TaskDoc(id string) string {
	return fmt.Sprintf("%s/%s", TasksDB(), id)
}

func NewTaskService(db FirestoreService[model.RecurringTask]) TaskService {
	return &taskServiceImpl{db: db}
}

func (service *taskServiceImpl) Paginate(start int, limit int) ([]model.RecurringTask, error) {
	query := service.db.Collection(TasksDB()).OrderBy(model.FieldCreatedAt, firestore.Desc)
	return service.db.Paginate(query, start, limit)
}
