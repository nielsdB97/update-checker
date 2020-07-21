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
	"path/filepath"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/joho/godotenv/autoload"
)

const hashesDir = "/hashes/"

func main() {
	fetchURL, parseURLErr := url.Parse(os.Getenv("URL"))
	checkErr(parseURLErr)
	_, fileName := path.Split(fetchURL.Path)

	ex, exErr := os.Executable()
	checkErr(exErr)
	workingDir := filepath.Dir(ex)
	hashesPath := workingDir + hashesDir

	fileContent, readFileErr := ioutil.ReadFile(hashesPath + fileName)

	if os.IsNotExist(readFileErr) {
		fmt.Println("File does not exist yet")

		if _, readHashesDirErr := os.Stat(hashesPath); os.IsNotExist(readHashesDirErr) {
			fmt.Println("Creating hashes directory")
			os.Mkdir(hashesPath, os.ModePerm)
		}
	}

	resp, httpErr := http.Get(fetchURL.String())
	checkErr(httpErr)

	defer resp.Body.Close()
	body, readBodyErr := ioutil.ReadAll(resp.Body)
	checkErr(readBodyErr)

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
	writeErr := ioutil.WriteFile(workingDir+hashesDir+fileName, bodyHashBytes, 0644)
	checkErr(writeErr)

	chatID, parseErr := strconv.ParseInt(os.Getenv("TG_CHAT_ID"), 10, 64)
	checkErr(parseErr)

	if readFileErr == nil {
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
