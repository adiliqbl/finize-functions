package service

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data/model"
	"fmt"
	"sync"
)

type TaskService interface {
	Doc(id string) *firestore.DocumentRef
	Paginate(start int, limit int) ([]model.RecurringTask, error)
	FindByUser(user string) ([]model.RecurringTask, error)
	FindByAccount(account string) ([]model.RecurringTask, error)
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

func (service *taskServiceImpl) Doc(id string) *firestore.DocumentRef {
	return service.db.Doc(TaskDoc(id))
}

func (service *taskServiceImpl) Paginate(start int, limit int) ([]model.RecurringTask, error) {
	query := service.db.Collection(TasksDB()).OrderBy(model.FieldCreatedAt, firestore.Desc)
	return service.db.Paginate(query, start, limit)
}

func (service *taskServiceImpl) FindByUser(user string) ([]model.RecurringTask, error) {
	return service.db.RunQuery(service.db.Collection(TasksDB()).Where(model.FieldUser, "==", user))
}

func (service *taskServiceImpl) FindByAccount(account string) (tasks []model.RecurringTask, err error) {
	var wg = &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		query := service.db.Collection(TasksDB()).Where(model.FieldData+"."+model.FieldAccountTo, "==", account)

		results, err1 := service.db.RunQuery(query)
		if err1 != nil {
			err = err1
		}
		tasks = append(tasks, results...)
	}()
	go func() {
		defer wg.Done()

		query := service.db.Collection(TasksDB()).Where(model.FieldData+"."+model.FieldAccountFrom, "==", account)

		results, err1 := service.db.RunQuery(query)
		if err1 != nil {
			err = err1
		}
		tasks = append(tasks, results...)
	}()

	wg.Wait()

	return tasks, err
}
