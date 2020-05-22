package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	fileContentHash := sha1.New()
	fileContentHash.Write(fileContent)
	fileContentHashBS := fileContentHash.Sum(nil)

	resp, httpErr := http.Get(ls50url)

	if httpErr != nil {
		panic(httpErr)
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		panic(readErr)
	}

	bodyHash := sha1.New()
	bodyHash.Write(body)
	bodyHashBS := bodyHash.Sum(nil)

	if bytes.Equal(fileContentHashBS, bodyHashBS) {
		fmt.Println("Files match!")
		return
	}

	fmt.Println("Writing file")
	writeErr := ioutil.WriteFile("/tmp/"+ls50FileName, body, 0644)

	if writeErr != nil {
		panic(writeErr)
	}
}
