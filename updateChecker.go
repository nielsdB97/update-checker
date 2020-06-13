package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fetchURL, parseURLErr := url.Parse(os.Getenv("URL"))
	checkErr(parseURLErr)
	_, fileName := path.Split(fetchURL.Path)

	fileContent, readErr := ioutil.ReadFile("/tmp/" + fileName)

	if readErr != nil {
		if !strings.Contains(readErr.Error(), "no such file or directory") {
			panic(readErr)
		}
		fmt.Println("File does not exist yet")
	}

	resp, httpErr := http.Get(fetchURL.String())
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
	writeErr := ioutil.WriteFile("/tmp/"+fileName, bodyHashBytes, 0644)
	checkErr(writeErr)

	chatID, parseErr := strconv.ParseInt(os.Getenv("TG_CHAT_ID"), 10, 64)
	checkErr(parseErr)

	if readErr != nil {
		fmt.Println("Sending notification")
		sendNotification(chatID, fetchURL.String())
	}
}

func sendNotification(chatID int64, updatedURL string) {
	bot, botErr := tgbotapi.NewBotAPI(os.Getenv("TG_API_TOKEN"))
	if botErr != nil {
		panic(botErr)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	messageString := fmt.Sprintf(`An update is available!
See: %s`, updatedURL)
	msg := tgbotapi.NewMessage(chatID, messageString)

	bot.Send(msg)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
