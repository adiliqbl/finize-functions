package model

import (
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapToEvent(t *testing.T) {
	got := model.Event{Processed: true}
	want, _ := util.MapTo[model.Event](map[string]interface{}{
		"processed": got.Processed,
	})

	assert.Equal(t, want, got)
}

func TestEventToMap(t *testing.T) {
	got := model.Event{Processed: true}
	want, _ := util.MapTo[map[string]interface{}](got)

	assert.Equal(t, true, want["processed"].(bool))
}
