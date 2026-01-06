package models

import "time"

type Event struct {
	ID          int
	Name        string    `bindings:"required"`
	Description string    `bindings:"required"`
	Location    string    `bindings:"required"`
	DateTime    time.Time `bindings:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	// later add it to a database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
