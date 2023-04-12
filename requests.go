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

	fmt.Println(token.Access_Token)
	bearer := "Bearer" + token.Token_Type
	if err != nil {
		log.Fatalln(err)
	}

	request.Header.Add("Authorization", bearer)
	request.Header.Add("User-Agent", "MyAPI/0.0.1")

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
