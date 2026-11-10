package models

import (
	"github.com/TuralMuzafarov/EventDiscovery/db"
	"github.com/TuralMuzafarov/EventDiscovery/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userID, err := result.LastInsertId()
	u.ID = userID
	return nil
}
