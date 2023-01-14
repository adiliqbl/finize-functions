package model

import (
	"finize-functions.app/data/model"
	"finize-functions.app/tests/fakedata"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserEventMapParsing(t *testing.T) {
	want := fakedata.NewUserEvent("id", "name", "email")
	got, _ := util.MapTo[model.UserEvent](fakedata.NewUserEventMap(want))

	assert.Equal(t, want, got)
}

func TestUserFromEvent(t *testing.T) {
	got := fakedata.NewUserEvent("id", "name", "email")
	want, _ := util.MapTo[model.User](got)

	assert.Equal(t, "id", want.ID)
	assert.Equal(t, "name", want.Name)
	assert.Equal(t, "email", want.Email)
}
