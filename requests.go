package main

import (
	"bytes"
	"fmt"
	"log"

	"net/http"
)

func Requests(token Token) {

	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, "https://oauth.reddit.com/r/AnimalTracking/new", nil)

	if err != nil {
		log.Fatalln(err)
	}

	request.Header.Add("Authorization", "Bearer "+token.Access_Token)
	request.Header.Add("User-Agent", "MyAPI/0.0.1")
	request.Header.Add("Content-Type", "application/json")

	log.Println("Setting headers...")

	response, err := client.Do(request)


	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Sending request...")

	if response.StatusCode != http.StatusOK {
		Error("Error")
	}

	defer response.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	data := buf.Bytes()

	//json.Unmarshal([]byte(data), &sub)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
}
