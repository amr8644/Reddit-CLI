package main

import (
	"encoding/json"
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Config struct {
	Client_ID string `yaml:"client_id"`
	Secret_ID string `yaml:"secret_id"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

type Token struct {
	Access_Token string `json:"access_token"`
	Token_Type   string `json:"token_type"`
	Expires      int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

func (c *Config) GetConfig() *Config {
	yaml_file, err := ioutil.ReadFile("config.yaml")

	Info("Reading config file...")

	if err != nil {
		Error(err.Error())
	}

	err = yaml.Unmarshal(yaml_file, c)

	if err != nil {
		Error(err.Error())
	}

	Success("Recieved credentials")

	return c

}

func Authenticate() Token {

	var c Config

	client := http.Client{}
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

	Info("Setting headers...")
	Info("Requesting access token")

	response, err := client.Do(request)

	if err != nil {
		Error(err.Error())
	}

	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)

	var token Token
	json.Unmarshal([]byte(result), &token)

	Success("Recived Access token")
	return token
}
