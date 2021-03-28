package handlers

import (
	model "github.com/tom-rt/jobless-bot/model"
	tb "gopkg.in/tucnak/telebot.v2"
)

//HandleIncomingMessage Function that handles every incoming messages
func HandleIncomingMessage(m *tb.Message) {
	var name = getName(m.Sender)
	// Checking user exists
	var userExists bool = userNameExists(name)
	// If doesn't, create it
	if (userExists) {
		model.AddMessage(name)
	} else { // User exists, incrementing messages count
		model.CreateUser(name)
	}
}

func userNameExists(name string) bool {
	_, err := model.GetUserByName(name)
	if err != nil {
		return false
	}
	return true
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