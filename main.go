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

	var report string  = handler.CreateReport()
	fmt.Println(report)

	c := cron.New()
	c.AddFunc("@every 1m", func() {
		fmt.Println("Every min")
		var report string  = handler.CreateReport()
		fmt.Println(report)
		// b.Send(CONV, "Ioos = beau goos")
	})
	c.Start()

	c2 := cron.New()
	c2.AddFunc("0 0 0 * * *", func() { fmt.Println("Every min 2") })
	c2.Start()

	c3 := cron.New()
	c3.AddFunc("* * * * * *", func() { fmt.Println("Every min 3") })
	c3.Start()

	// c4 := cron.New()
	// c4.AddFunc("* * * * *", func() { fmt.Println("Every min 4") })
	// c4.Start()

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

	b.Handle("/au_revoir", func(m *tb.Message) {
		b.Send(m.Chat, "Au revoir l'élite")
	})

	// b.Handle("/stats", func(m *tb.Message) {
	// 	report := handler.CreateReport(m)
	// 	b.Send(m.Chat, report)
	// })

	b.Start()
}
