package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEvent(t *testing.T) {
	err := eventService.SetProcessed(nil)
	assert.Nil(t, err)

	processed := eventService.IsProcessed()
	assert.True(t, processed)
}
