package main

import (
	"flag"
	"net/http"
)

type BaseClient struct {
	cl *http.Client
}

func NewClient() *BaseClient{
    return &BaseClient{
        cl:&http.Client{},
    }
}

func main() {

    client := NewClient()

	username := flag.String("u", "username", "Username you want search")
    where := flag.String("w","about", "The information you want to get")
	flag.Parse()

	token := client.Authenticate()
	client.UserInformation(token, *username,*where)

}
