package main

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] (chatID: %s) %s", update.Message.From.UserName, strconv.FormatInt(update.Message.Chat.ID, 10), update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Message received!")
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
