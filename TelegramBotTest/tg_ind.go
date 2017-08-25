package main

import (
	"flag"
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Water"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Electricity"),
	),
)

func main() {
	token := flag.String("token", "", "Token for telegram API")

	bot, err := tgbotapi.NewBotAPI(*token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "/new":
			msg.ReplyMarkup = numericKeyboard
		case "/finish":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		bot.Send(msg)
	}
}
