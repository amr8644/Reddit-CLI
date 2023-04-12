package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reddit/cli/types"
)

var base_url string = "https://oauth.reddit.com/"

func Requests(token Token, username string) {

	var ty types.UserAbout
	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, base_url+"user/"+username+"/about", nil)

	bearer := "Bearer " + token.Access_Token

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

	data, err := ioutil.ReadAll(response.Body)

	json.Unmarshal([]byte(data), &ty)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(ty)
}

func UsersAbout(token Token,username string) {

	Requests(token,username)
}
