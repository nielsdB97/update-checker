package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	urlToCheck := "https://assets.kef.com/pdf_doc/ls50w/LS50-Wireless-Firmware-Release-Note.pdf"
	resp, httpErr := http.Get(urlToCheck)

	if httpErr != nil {
		panic(httpErr)
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		panic(readErr)
	}

	writeErr := ioutil.WriteFile("/tmp/ls50_release_notes", body, 0644)

	if writeErr != nil {
		panic(writeErr)
	}
}
