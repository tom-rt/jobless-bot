package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	b.Handle(tb.OnText, func(m *tb.Message) {
		var name = getName(m.Sender)
		fmt.Println(name + " dit: " + m.Text)
	})

	b.Handle("/salut", func(m *tb.Message) {
		b.Send(m.Chat, "Salut l'élite")
	})

	b.Start()
}

func getName(sender *tb.User) string {
	if sender.FirstName == "" {
		return sender.Username
	} else if sender.LastName != "" {
		return sender.FirstName + " " + sender.LastName
	} else if sender.Username == "" {
		return sender.FirstName
	} else {
		return sender.FirstName + " " + sender.Username
	}
}