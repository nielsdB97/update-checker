package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	bot, botErr := tgbotapi.NewBotAPI(os.Getenv("TG_API_TOKEN"))
	if botErr != nil {
		panic(botErr)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, updatesErr := bot.GetUpdatesChan(u)
	if updatesErr != nil {
		panic(updatesErr)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] (chatID: %s) %s", update.Message.From.UserName, string(update.Message.Chat.ID), update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Message received!")
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
