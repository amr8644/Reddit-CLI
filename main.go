package main

import (
	"encoding/json"
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (c *Config) GetConfig() *Config {
	yaml_file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(yaml_file, c)

	if err != nil {
		log.Fatalln(err)
	}

	return c
}

func GetAccessToken(client *http.Client) Token {

	var c Config
	c.GetConfig()

	data := url.Values{}
	data.Set("grant_type", "password")

	data.Set("username", c.Username)
	data.Set("password", c.Password)
	data.Set("client_id", c.Client_ID)
	data.Set("secret_id", c.Secret_ID)

	request, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader(data.Encode()))

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

    client.GetTopPosts()
	fmt.Println(token)

}
