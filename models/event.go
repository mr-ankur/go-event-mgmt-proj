package models

import (
	"time"

	"example.com/event-mgmt/db"
)

type Event struct {
	ID          int64
	Name        string    `bindings:"required"`
	Description string    `bindings:"required"`
	Location    string    `bindings:"required"`
	DateTime    time.Time `bindings:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	// later add it to a database
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func GetAllEvents() []Event {
	return events
}
