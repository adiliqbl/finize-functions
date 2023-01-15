package model

import (
	"finize-functions.app/data/model"
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserEventMapParsing(t *testing.T) {
	want := fake.NewUserEvent("id", "name", "email")
	got, _ := util.MapTo[model.UserEvent](fake.NewUserEventMap(want))

	assert.Equal(t, want, got)
}

func TestUserFromEvent(t *testing.T) {
	got := fake.NewUserEvent("id", "name", "email")
	want, _ := util.MapTo[model.User](got)

	assert.Equal(t, "id", want.ID)
	assert.Equal(t, "name", want.Name)
	assert.Equal(t, "email", want.Email)
}
