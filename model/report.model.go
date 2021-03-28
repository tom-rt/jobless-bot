package models

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx/types"
	db "github.com/tom-rt/jobless-bot/db"
)

// User model
type Report struct {
	TotalMessageCount	int	`db:"total_message_count", json:"totalMessageCount" binding:"required"`
	UsersReports		types.JSONText	`db:"users_reports" json:"usersReports" binding:"required"`
}

// GetReport generates a report on sent messages
func GetReport() (Report, error) {
	var report Report
	var spammers []string
	var reportMessage string = "Salut l'élite !\n"

	err := db.DB.Select(&spammers,
		`SELECT name FROM chan_user WHERE sent_messages_count = (SELECT MAX(sent_messages_count) FROM chan_user)`,
	)

	fmt.Println(spammers)

	err = db.DB.Get(&report,
		`
		SELECT	SUM(sent_messages_count) AS total_message_count,
		(
			SELECT json_agg(json_build_object('name', name, 'sent_messages_count', sent_messages_count) ORDER BY sent_messages_count DESC) FROM chan_user
		) AS users_reports
		FROM chan_user;
		`)

	if len(spammers) > 1 {
		reportMessage = "Félicitations aux meilleurs chomeurs des dernières 24 heures:"
		for i := 0; i < len(spammers); i++ {
			if i == len(spammers) - 1 {
				reportMessage = reportMessage + " et " + spammers[i] + " avec un total de " + strconv.Itoa(report.TotalMessageCount) + " messages chacun !"
			} else {
				reportMessage = reportMessage +" " + spammers[i] + ","
			}
		}
	}

	return report, err
}
