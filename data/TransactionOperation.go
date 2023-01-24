package data

import (
	"cloud.google.com/go/firestore"
)

type TransactionOperation struct {
	Ref    *firestore.DocumentRef
	Data   interface{}
	Create bool
}

func Commit(tx *firestore.Transaction, ops []TransactionOperation) error {
	for _, operation := range ops {
		var err error
		if _, ok := operation.Data.([]firestore.Update); ok {
			err = tx.Update(operation.Ref, operation.Data.([]firestore.Update))
		} else if operation.Create {
			err = tx.Create(operation.Ref, operation.Data)
		} else {
			err = tx.Set(operation.Ref, operation.Data)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func Perform(batch *firestore.BulkWriter, ops []TransactionOperation) error {
	for _, operation := range ops {
		var err error
		if _, ok := operation.Data.([]firestore.Update); ok {
			_, err = batch.Update(operation.Ref, operation.Data.([]firestore.Update))
		} else if operation.Create {
			_, err = batch.Create(operation.Ref, operation.Data)
		} else {
			_, err = batch.Set(operation.Ref, operation.Data)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
