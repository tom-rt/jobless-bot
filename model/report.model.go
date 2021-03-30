package models

import (
	"encoding/json"

	"github.com/jmoiron/sqlx/types"
	db "github.com/tom-rt/jobless-bot/db"
)

// Report query model
type ReportQuery struct {
	TotalMessageCount	int				`db:"totalMessageCount" json:"totalMessageCount" binding:"required"`
	UsersReports		types.JSONText	`db:"usersReports" json:"usersReports" binding:"required"`
}

// Report query
type Report struct {
	TotalMessageCount	int		`json:"totalMessageCount" binding:"required"`
	UsersReports		[]User	`json:"usersReports" binding:"required"`
}

// GetReport generates a report on sent messages
func GetReport() (*Report, []string, int, error) {
	var query ReportQuery
	var maxCount int
	var spammers []string

	err := db.DB.Select(&spammers,
		`SELECT name FROM chan_user WHERE sent_messages_count = (SELECT MAX(sent_messages_count) FROM chan_user)`,
	)

	err = db.DB.Get(&maxCount,
		`SELECT MAX(sent_messages_count) FROM chan_user`,
	)

	err = db.DB.Get(&query,
		`SELECT	SUM(sent_messages_count) AS "totalMessageCount",
		(
			SELECT json_agg(json_build_object('id', id, 'name', name, 'sentMessagesCount', sent_messages_count) ORDER BY sent_messages_count DESC) FROM chan_user
		) AS "usersReports"
		FROM chan_user;
	`)

	var report *Report = new(Report)
	report.TotalMessageCount = query.TotalMessageCount
	json.Unmarshal(query.UsersReports, &report.UsersReports)

	return report, spammers, maxCount, err
}

func ResetReport() {
	tx := db.DB.MustBegin()
	tx.MustExec("DELETE FROM chan_user WHERE sent_messages_count = 0;")
	tx.MustExec("UPDATE chan_user SET sent_messages_count = 0;")
	tx.Commit()
}