package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetTopPosts(token Token) {

	client := http.Client{}

	request, err := http.NewRequest("GET", "https://oauth.reddit.com/user/genericlemon24/overview", nil)

    var user Redditor
	if err != nil {
		log.Fatalln(err)

	}

	request.Header.Set("Authorization", "bearer"+token.Access_Token)
	request.Header.Set("User-Agent", "ChangeMeClient/0.1 by YourUsername")
	request.Header.Set("Content-Type", "application/json")
	
    log.Println("Setting headers...")

	response, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Sending request...")

	defer response.Body.Close()

	//	result, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	json.NewDecoder(response.Body).Decode(&user)

    fmt.Println(user)

}
