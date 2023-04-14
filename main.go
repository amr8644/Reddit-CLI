package main

import (
	"flag"
	"net/http"
	"os"
)

type BaseClient struct {
	cl *http.Client
}

func NewClient() *BaseClient {
	return &BaseClient{
		cl: &http.Client{},
	}
}

func main() {

	client := NewClient()

	os.Mkdir("output",0600)
	username := flag.String("u", "username", "Username you want search")
//	where := flag.String("w", "about", "The information you want to get")
	flag.Parse()

	token := client.Authenticate()

	client.UserUpvotes(token, *username)

}
