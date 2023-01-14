package tests

import (
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/tests/fakedata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsingDocumentPath(t *testing.T) {
	// Simple
	event := data.FirestoreEvent[model.UserEvent]{
		OldValue: data.FirestoreValue[model.UserEvent]{
			Name: "projects/projectId/databases/(default)/documents/users/userId",
			Data: fakedata.NewUserEvent("id", "name", "email"),
		},
		Value: data.FirestoreValue[model.UserEvent]{
			Name: "projects/projectId/databases/(default)/documents/users/userId",
			Data: fakedata.NewUserEvent("id", "new name", "new email"),
		},
	}

	path, collection, doc := event.Path()

	assert.Equal(t, "users/userId", path)
	assert.Equal(t, "users", collection)
	assert.Equal(t, "userId", doc)

	// Complex
	event = data.FirestoreEvent[model.UserEvent]{
		OldValue: data.FirestoreValue[model.UserEvent]{
			Name: "projects/projectId/databases/(default)/documents/user-transactions/userId/transactions/transactionId",
			Data: fakedata.NewUserEvent("id", "name", "email"),
		},
		Value: data.FirestoreValue[model.UserEvent]{
			Name: "projects/projectId/databases/(default)/documents/user-transactions/userId/transactions/transactionId",
			Data: fakedata.NewUserEvent("id", "new name", "new email"),
		},
	}

	path, collection, doc = event.Path()

	assert.Equal(t, "user-transactions/userId/transactions/transactionId", path)
	assert.Equal(t, "user-transactions", collection)
	assert.Equal(t, "userId/transactions/transactionId", doc)
}
