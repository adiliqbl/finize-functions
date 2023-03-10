package data

import (
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/tests/fake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsingDocumentPath(t *testing.T) {
	// Simple
	event := data.FirestoreEvent[model.UserEvent]{
		OldValue: data.FirestoreValue[model.UserEvent]{
			Name: "projects/projectId/databases/(default)/documents/users/userId",
			Data: fake.NewUserEvent("id", "name", "email"),
		},
		Value: data.FirestoreValue[model.UserEvent]{
			Name: "projects/projectId/databases/(default)/documents/users/userId",
			Data: fake.NewUserEvent("id", "new name", "new email"),
		},
	}

	path, collection, doc := event.Path()

	assert.Equal(t, "users/userId", path)
	assert.Equal(t, "users", collection)
	assert.Equal(t, "userId", doc)
	assert.Equal(t, "userId", event.UserID())

	// Complex
	event = data.FirestoreEvent[model.UserEvent]{
		OldValue: data.FirestoreValue[model.UserEvent]{
			Name: "projects/projectId/databases/(default)/documents/user-transactions/userId/transactions/transactionId",
			Data: fake.NewUserEvent("id", "name", "email"),
		},
		Value: data.FirestoreValue[model.UserEvent]{
			Name: "projects/projectId/databases/(default)/documents/user-transactions/userId/transactions/transactionId",
			Data: fake.NewUserEvent("id", "new name", "new email"),
		},
	}

	path, collection, doc = event.Path()

	assert.Equal(t, "user-transactions/userId/transactions/transactionId", path)
	assert.Equal(t, "user-transactions", collection)
	assert.Equal(t, "userId", event.UserID())
	assert.Equal(t, "userId/transactions/transactionId", doc)
}
