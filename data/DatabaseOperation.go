package data

import (
	"cloud.google.com/go/firestore"
)

type DatabaseOperation struct {
	Ref    *firestore.DocumentRef
	Data   interface{}
	Create bool
}

func CommitTransaction(tx *firestore.Transaction, ops []DatabaseOperation) error {
	for _, operation := range ops {
		var err error
		if _, ok := operation.Data.([]firestore.Update); ok {
			err = tx.Update(operation.Ref, operation.Data.([]firestore.Update))
		} else if operation.Create {
			err = tx.Create(operation.Ref, operation.Data)
		} else {
			err = tx.Set(operation.Ref, operation.Data, firestore.MergeAll)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func CommitBatch(batch *firestore.BulkWriter, ops []DatabaseOperation) error {
	for _, operation := range ops {
		var err error
		if updates, ok := operation.Data.([]firestore.Update); ok {
			_, err = batch.Update(operation.Ref, updates)
		} else if update, ok := operation.Data.(firestore.Update); ok {
			_, err = batch.Update(operation.Ref, []firestore.Update{update})
		} else if operation.Create {
			_, err = batch.Create(operation.Ref, operation.Data)
		} else {
			_, err = batch.Set(operation.Ref, operation.Data, firestore.MergeAll)
		}

		if err != nil {
			return err
		}
	}

	batch.Flush()
	return nil
}
