package models

import (
	"example.com/event-mgmt/db"
	"example.com/event-mgmt/utils"
)

type User struct {
	ID       int64
	Email    string `bindings:"required"`
	Password string `bindings:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users(email, password)
		VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id

	return err
}
