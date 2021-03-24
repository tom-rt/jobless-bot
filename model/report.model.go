package models

import (
	db "github.com/tom-rt/jobless-bot/db"
)

// User model
type Report struct {
	TotalMessageCount	int64	`db:"total_message_count", json:"totalMessageCount" binding:"required"`
	UsersReports		[]User	`db:"users_reports" json:"usersReports" binding:"required"`
}

// GetReport generates a report on sent messages
func GetReport() ([]Report, error) {
	var report []Report
	err := db.DB.Select(&report,
		`
		SELECT	SUM(sent_messages_count) AS total_message_count,
				(
					SELECT name, sent_messages_count FROM chan_user AS users_reports ORDER BY sent_messages_count DESC
				)
		FROM chan_user
		`)
	return report, err
}