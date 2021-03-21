package main

import (
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	db "github.com/tom-rt/jobless-bot/db"
	handler "github.com/tom-rt/jobless-bot/handler"

	tb "gopkg.in/tucnak/telebot.v2"
)

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

	b.Handle(tb.OnText, func(m *tb.Message) {
		handler.HandleIncomingMessage(m)
	})

	b.Handle("/salut", func(m *tb.Message) {
		b.Send(m.Chat, "Salut l'élite")
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
