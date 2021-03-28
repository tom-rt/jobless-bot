package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	db "github.com/tom-rt/jobless-bot/db"
	handler "github.com/tom-rt/jobless-bot/handler"

	cron "github.com/robfig/cron/v3"
	tb "gopkg.in/tucnak/telebot.v2"
)

var CONV *tb.Chat = nil

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token: os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	db.InitDB()

	var tmpReport string  = handler.CreateReport()
	fmt.Println(tmpReport)

	reprise := cron.New()
	// Every monday at 7 AM
	reprise.AddFunc("0 0 7 * * 1", func() {
		link := handler.SendReprise()
		b.Send(CONV, link)
	})
	reprise.Start()

	cronReport := cron.New()
	// Every day at 10 AM
	cronReport.AddFunc("0 0 10 * * *", func() {
		report := handler.CreateReport()
		fmt.Println(report)
		// b.Send(CONV, report)
	})
	cronReport.Start()

	b.Handle(tb.OnText, func(m *tb.Message) {
		if CONV == nil {
			CONV = m.Chat
		}
		handler.HandleIncomingMessage(m)
	})

	b.Handle("/salut", func(m *tb.Message) {
		b.Send(m.Chat, "Salut l'élite")
	})

	b.Handle("/stats", func(m *tb.Message) {
		var report string  = handler.CreateReport()
		fmt.Println(report)
		// b.Send(m.Chat, "Salut l'élite")
	})

	b.Handle("/bonne_nuit", func(m *tb.Message) {
		b.Send(m.Chat, "Bonne nuit l'élite")
	})

	b.Start()
}
