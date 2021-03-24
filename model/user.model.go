package models

import (
	db "github.com/tom-rt/jobless-bot/db"
)

// User model
type User struct {
	ID					int		`db:"id" json:"id"`
	Name				string	`db:"name" json:"name" binding:"required"`
	SentMessagesCount	string	`db:"sent_messages_count" json:"sent_messages_count" binding:"required"`
}

// GetUserByName func
func GetUserByName(name string) (User, error) {
	var user User
	err := db.DB.Get(&user,
		`SELECT		id,
					name,
					sent_messages_count
		FROM chan_user
		WHERE name = $1`,
		name)
	return user, err
}

// CreateUser func
func CreateUser(name string) error {
	var user User
	err := db.DB.Get(&user,
		`INSERT INTO chan_user
			(name)
			VALUES
			($1)
		`,
		name)
	return err
}

// AddMessage increments sent messages count for a given user
func AddMessage(name string) {
	tx := db.DB.MustBegin()
	tx.MustExec("UPDATE chan_user SET sent_messages_count = sent_messages_count + 1 WHERE name = $1", name)
	tx.Commit()
}