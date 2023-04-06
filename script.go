package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Token struct {
	Access_Token string `json:"access_token"`
	Token_Type   string `json:"token_type"`
	Expires      int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

func GetAccessToken(client *http.Client) Token {

	data := url.Values{}
	data.Set("grant_type", "password")

	data.Set("client_id", CLIENT_ID)
	data.Set("secret_id", SECRET_ID)

	request, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader(data.Encode()))

	request.SetBasicAuth(CLIENT_ID, SECRET_ID)
	request.Header.Set("User-Agent", "MyAPI/0.0.1")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)

	var token Token
	json.Unmarshal([]byte(result), &token)
	log.Println("Got the access token")

	return token
}

func main() {

	client := &http.Client{}
	token := GetAccessToken(client)

	request, err := http.NewRequest("GET", "https://oauth.reddit.com/api/v1/me", nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Add("Authorization", "bearer"+token.Access_Token)
	request.Header.Set("User-Agent", "ChangeMeClient/0.1 by YourUsername")
	log.Println("Setting headers...")
	response, err := client.Do(request)

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
