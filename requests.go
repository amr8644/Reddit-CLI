package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

var base_url string = "https://oauth.reddit.com/"

func (client *BaseClient) Requests(token Token, url string) ([]byte, error) {

	request, err := http.NewRequest(http.MethodGet, base_url+url, nil)

	bearer := "Bearer " + token.Access_Token

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	request.Header.Add("Authorization", bearer)
	request.Header.Add("User-Agent", "MyAPI/0.0.1")

	response, err := client.cl.Do(request)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return data, nil
}
