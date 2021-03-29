package models

import (
	"fmt"

	"github.com/jmoiron/sqlx/types"
	db "github.com/tom-rt/jobless-bot/db"
)

// User model
type Report struct {
	TotalMessageCount	int				`db:"total_message_count", json:"totalMessageCount" binding:"required"`
	UsersReports		types.JSONText	`db:"users_reports" json:"usersReports" binding:"required"`
}

// GetReport generates a report on sent messages
func GetReport() (Report, []string, int, error) {
	var report Report
	var maxCount int
	var spammers []string

	err := db.DB.Select(&spammers,
		`SELECT name FROM chan_user WHERE sent_messages_count = (SELECT MAX(sent_messages_count) FROM chan_user)`,
	)

	err = db.DB.Get(&maxCount,
		`SELECT MAX(sent_messages_count) FROM chan_user`,
	)

	err = db.DB.Get(&report,
		`
		SELECT	SUM(sent_messages_count) AS total_message_count,
		(
			SELECT json_agg(json_build_object('name', name, 'sent_messages_count', sent_messages_count) ORDER BY sent_messages_count DESC) FROM chan_user
		) AS users_reports
		FROM chan_user;
		`)

	fmt.Println(spammers)

	return report, spammers, maxCount, err
}
