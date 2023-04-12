package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	yaml "gopkg.in/yaml.v3"
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

	if err != nil {
		Error(err.Error())
	}

	err = yaml.Unmarshal(yaml_file, c)

	if err != nil {
		Error(err.Error())
	}

	return c

}

func Authenticate() Token {

	var c Config

	client := http.Client{}
	credentials := c.GetConfig()

	data := url.Values{}

	data.Set("grant_type", "password")
	data.Set("username", credentials.Username)
	data.Set("password", credentials.Password)

	request, err := http.NewRequest(http.MethodPost, "https://www.reddit.com/api/v1/access_token", strings.NewReader(data.Encode()))

	request.SetBasicAuth(credentials.Client_ID, credentials.Secret_ID)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("User-Agent", "MyAPI/0.0.1")
	response, err := client.Do(request)

	if err != nil {
		Error(err.Error())
	}

	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	var token Token
	json.Unmarshal([]byte(string(result)), &token)

	return token
}
