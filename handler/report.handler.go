package handlers

import (
	"fmt"
	"strconv"

	model "github.com/tom-rt/jobless-bot/model"
)

// CreateReport returns the stats
func CreateReport() string {
	report, spammers, maxCount, err := model.GetReport()
	fmt.Println(len(spammers), spammers)

	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("report", report)
	}

	var reportMessage string = "Salut l'élite !\n"
	if len(spammers) > 1 {
		reportMessage = reportMessage + "Félicitations aux meilleurs chomeurs des dernières 24 heures:"
		for i := 0; i < len(spammers); i++ {
			if i == len(spammers) - 1 {
				reportMessage = reportMessage + " et " + spammers[i] + " avec un total de " + strconv.Itoa(maxCount) + " messages chacun !"
			} else {
				reportMessage = reportMessage +" " + spammers[i] + ","
			}
		}
	} else if len(spammers) == 1 {
		reportMessage = reportMessage + "Félicitations au meilleur chomeur des dernières 24 heures: " + spammers[0] + " avec un total de " + strconv.Itoa(maxCount) + " messages !"
	} else {
		return "Personne n'a parlé !"
	}


	reportMessage = reportMessage + "\n\nVoici le classement:\n"

	for i := 0; i < len(report.UsersReports); i++ {
		fmt.Println(report.UsersReports[i])
		reportMessage = reportMessage + "\n - " + report.UsersReports[i].Name + ": " + report.UsersReports[i].SentMessagesCount + "."
	}

	fmt.Println(reportMessage)

	return "stringReport"
}

func SendReprise() string {
	return "https://www.youtube.com/watch?v=3L5N4qypyyY"
}