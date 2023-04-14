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

	os.Mkdir("output", 0777)
	username := flag.String("u", "username", "Username you want search")
	// subreddit := flag.String("s","subreddit","Subreddit you want to search")
	//	where := flag.String("w", "about", "The information you want to get")
	//	limit := flag.String("l", "25", "the maximum number of items desired (default: 25, maximum: 100)")
	//	sort := flag.String("s", "hot", "one of (hot, new, top, controversial)")
	//	time := flag.String("t", "week", "one of (hour, day, week, month, year, all)")
	flag.Parse()

	token := client.Authenticate()

	client.UserAbout(token, *username)

}
