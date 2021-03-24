package handlers

import (
	"fmt"

	model "github.com/tom-rt/jobless-bot/model"
)

// CreateReport returns the stats
func CreateReport() string {
	report, err := model.GetReport()
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("report", report)
	}
	// stringReport := "Félicitations au meilleur chômeur des dernières 24 heures: " + report[0].Name + " !\n"
	return "stringReport"
}