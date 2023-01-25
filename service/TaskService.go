package service

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data/model"
	"finize-functions.app/util"
)

type TaskService interface {
	Create(task model.RecurringTask) (*string, error)
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

func (service *taskServiceImpl) Create(task model.RecurringTask) (*string, error) {
	data, _ := util.MapTo[map[string]interface{}](task)
	return service.db.Create(tasksDB(), nil, data)
}

func (service *taskServiceImpl) PaginateTasks(start uint, limit uint) ([]model.RecurringTask, error) {
	query := service.db.Collection(tasksDB()).OrderBy("created_at", firestore.Desc)
	return service.db.Paginate(query, start, limit)
}
