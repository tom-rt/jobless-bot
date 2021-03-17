package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	tb "gopkg.in/tucnak/telebot.v2"
)

// DB golbal sqlx  connection instance
var DB *sqlx.DB

//InitDB func
func InitDB() {
	var err error
	var dbHost string = os.Getenv("DB_HOST")
	var dbPort string = os.Getenv("DB_PORT")
	var dbUser string = os.Getenv("DB_USER")
	var dbPassword string = os.Getenv("DB_PWD")
	var dbName string = os.Getenv("DB_NAME")

	//SQLX
	var dbConnection string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	// var dbConnection string = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbName)
	DB, err = sqlx.Connect("postgres", dbConnection)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		os.Exit(1)
	}
}

// CloseDB closes the connection of the current instance
func CloseDB() {
	DB.Close()
}


func main() {
	b, err := tb.NewBot(tb.Settings{
		Token: os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})


	if err != nil {
		log.Fatal(err)
		return
	}

	InitDB()

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
	if sender.Username != "" {
		return sender.Username
	} else if sender.FirstName != "" && sender.LastName != "" {
		return sender.FirstName + " " + sender.LastName
	} else if sender.FirstName != "" {
		return sender.FirstName
	} else {
		return sender.LastName
	}
}