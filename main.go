package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	db "github.com/tom-rt/jobless-bot/db"
	handler "github.com/tom-rt/jobless-bot/handler"

	gocron "github.com/go-co-op/gocron"
	tb "gopkg.in/tucnak/telebot.v2"
)

var CONV *tb.Chat = nil
var BOT *tb.Bot = nil
var LAST_MESSAGE string = ""

func main() {

	BOT, err := tb.NewBot(tb.Settings{
		Token: os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	db.InitDB()
	localTime, err := time.LoadLocation("Europe/Paris")

	var tmpReport string  = handler.CreateReport()
	fmt.Println(tmpReport)

	// TEST
	// tmps := gocron.NewScheduler(localTime)
	// tmps.Every(1).Day().Monday().At("19:32").Do(func () {BOT.Send(CONV, "Je suis immatériel, mais j'ai des sentiments.")})
	// tmps.StartAsync()

	// REPRISE EVERY MONDAY MORNING
	repriseScheduler := gocron.NewScheduler(localTime)
	repriseScheduler.Every(1).Day().Monday().At("7:00").Do(sendReprise)
	repriseScheduler.StartAsync()

	// REPORT EVERY MORNING AT 9
	reportScheduler := gocron.NewScheduler(localTime)
	reportScheduler.Every(1).Day().At("9:00").Do(sendReport)
	reportScheduler.StartAsync()


	BOT.Handle(tb.OnText, func(m *tb.Message) {
		if CONV == nil {
			CONV = m.Chat
		}
		handler.HandleIncomingMessage(m)

		if (m.Text == LAST_MESSAGE) {
			BOT.Send(m.Chat, "Chips")
		}
		LAST_MESSAGE = m.Text
	})

	BOT.Handle("/salut", func(m *tb.Message) {
		BOT.Send(m.Chat, "Salut l'élite")
	})

	BOT.Handle("/stats", func(m *tb.Message) {
		sendReport()
	})

	BOT.Handle("/bonne_nuit", func(m *tb.Message) {
		BOT.Send(m.Chat, "Bonne nuit l'élite")
	})

	BOT.Start()
}

func sendReprise() {
	handler.SendReprise()
	link := handler.SendReprise()
	BOT.Send(CONV, link)
}

func sendReport() {
	handler.SendReprise()
	report := handler.CreateReport()
	fmt.Println(report)
}