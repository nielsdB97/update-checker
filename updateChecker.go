package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/joho/godotenv/autoload"
)

const ls50FileName = "ls50_release_notes"
const ls50url = "https://assets.kef.com/pdf_doc/ls50w/LS50-Wireless-Firmware-Release-Note.pdf"

func main() {
	fileContent, readErr := ioutil.ReadFile("/tmp/" + ls50FileName)

	if readErr != nil {
		if !strings.Contains(readErr.Error(), "no such file or directory") {
			panic(readErr)
		}
		fmt.Println("File does not exist yet")
	}

	resp, httpErr := http.Get(ls50url)
	checkErr(httpErr)

	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	checkErr(readErr)

	bodyHash := sha1.New()
	bodyHash.Write(body)
	bodyHashBytes := bodyHash.Sum(nil)

	fmt.Printf("Previous: %x\n", fileContent)
	fmt.Printf("Incoming: %x\n", bodyHashBytes)
	if bytes.Equal(fileContent, bodyHashBytes) {
		fmt.Println("Hashes match!")
		return
	}

	fmt.Println("Writing file")
	writeErr := ioutil.WriteFile("/tmp/"+ls50FileName, bodyHashBytes, 0644)
	checkErr(writeErr)

	chatID, parseErr := strconv.ParseInt(os.Getenv("TG_CHAT_ID"), 10, 64)
	checkErr(parseErr)

	fmt.Println("Sending notification")
	sendNotification(chatID)
}

func sendNotification(chatID int64) {
	bot, botErr := tgbotapi.NewBotAPI(os.Getenv("TG_API_TOKEN"))
	if botErr != nil {
		panic(botErr)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	msg := tgbotapi.NewMessage(chatID, "An update is available for the KEF LS50 Wireless!")

	bot.Send(msg)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
