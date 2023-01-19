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
