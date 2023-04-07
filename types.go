package main




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

