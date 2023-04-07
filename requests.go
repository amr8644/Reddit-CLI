package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (client HTTP) GetTopPosts(token Token) {

	request, err := http.NewRequest("GET", "https://oauth.reddit.com/golang/top", nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Add("Authorization", "bearer"+token.Access_Token)
	request.Header.Set("User-Agent", "ChangeMeClient/0.1 by YourUsername")
	log.Println("Setting headers...")
	response, err := client.Client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Sending requests...")
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(result))

}
