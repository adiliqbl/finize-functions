package service

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"fmt"
	"log"
)

type EventService interface {
	IsProcessed() bool
	SetProcessed(tx *firestore.Transaction) error
	SetProcessedBatch(batch *firestore.BulkWriter) error
}

type eventServiceImpl struct {
	db      FirestoreService[model.Event]
	eventID string
}

func eventsDB() string {
	return "events"
}

func eventDoc(id string) string {
	return fmt.Sprintf("%s/%s", eventsDB(), id)
}

func NewEventService(db FirestoreService[model.Event], eventID string) EventService {
	return &eventServiceImpl{db: db, eventID: eventID}
}

func (service *eventServiceImpl) IsProcessed() bool {
	event, err := service.db.Find(eventDoc(service.eventID), nil)
	if err != nil {
		log.Printf("IsTransactionProcessed: %v", err)
		return false
	}
	return event.Processed
}

func (service *eventServiceImpl) SetProcessed(tx *firestore.Transaction) error {
	doc, err := util.MapTo[map[string]interface{}](model.Event{Processed: true})
	if err != nil {
		log.Fatalf("SetProcessed – Failed to convert event to map: %v", err)
		return err
	}

	if tx != nil {
		return tx.Set(service.db.Doc(eventDoc(service.eventID)), doc)
	} else {
		_, err = service.db.Create(eventsDB(), &service.eventID, doc)
		return err
	}
}

func (service *eventServiceImpl) SetProcessedBatch(batch *firestore.BulkWriter) error {
	doc, err := util.MapTo[map[string]interface{}](model.Event{Processed: true})
	if err != nil {
		log.Fatalf("SetProcessed – Failed to convert event to map: %v", err)
		return err
	}

	_, err = batch.Set(service.db.Doc(eventDoc(service.eventID)), doc)
	return err
}
