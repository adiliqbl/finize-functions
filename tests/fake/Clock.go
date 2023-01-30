package fake

import (
	"finize-functions.app/util"
	"log"
	"time"
)

func NewClock(year int, month time.Month, day int) util.Clock {
	return &fakeClock{year: year, month: month, day: day}
}

type fakeClock struct {
	year  int
	month time.Month
	day   int
}

func (clock fakeClock) Now() time.Time {
	return time.Date(clock.year, clock.month, clock.day, 0, 0, 0, 0, time.UTC).UTC()
}

func (clock fakeClock) WithZone(timezone string) (*time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatalf("Failed to parse timezone: %v", err)
		return nil, err
	}
	return util.Pointer(clock.Now().In(location)), nil
}
