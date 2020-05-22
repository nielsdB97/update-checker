package main

import (
	"io/ioutil"
	"net/http"
)

const ls50FileName = "ls50_release_notes"
const ls50url = "https://assets.kef.com/pdf_doc/ls50w/LS50-Wireless-Firmware-Release-Note.pdf"

func main() {
	resp, httpErr := http.Get(ls50url)

	if httpErr != nil {
		panic(httpErr)
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		panic(readErr)
	}

	writeErr := ioutil.WriteFile("/tmp/"+ls50FileName, body, 0644)

	if writeErr != nil {
		panic(writeErr)
	}
}
