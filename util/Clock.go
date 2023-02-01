package util

import (
	"log"
	"time"
)

type Clock interface {
	Now() time.Time
	WithZone(timezone string) (*time.Time, error)
}

func NewClock() Clock {
	return &realClock{}
}

type realClock struct{}

func (clock realClock) Now() time.Time { return time.Now().UTC() }

func (clock realClock) WithZone(timezone string) (*time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		log.Printf("Failed to parse timezone: %v", err)
		return nil, err
	}
	return Pointer(clock.Now().In(location)), nil
}
