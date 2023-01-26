package service

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data/model"
)

type TaskService interface {
	PaginateTasks(start uint, limit uint) ([]model.RecurringTask, error)
}

type taskServiceImpl struct {
	db FirestoreService[model.RecurringTask]
}

func tasksDB() string {
	return "tasks"
}

func NewTaskService(db FirestoreService[model.RecurringTask]) TaskService {
	return &taskServiceImpl{db: db}
}

func (service *taskServiceImpl) PaginateTasks(start uint, limit uint) ([]model.RecurringTask, error) {
	query := service.db.Collection(tasksDB()).OrderBy(model.FieldCreatedAt, firestore.Desc)
	return service.db.Paginate(query, start, limit)
}
