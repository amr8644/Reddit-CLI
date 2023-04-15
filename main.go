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
	where := flag.String("w", "about", "The information you want to get")
	limit := flag.String("limit", "25", "the maximum number of items desired (default: 25, maximum: 100)")
//	sort := flag.String("sort", "hot", "one of (hot, new, top, controversial)")
	time := flag.String("time", "week", "one of (hour, day, week, month, year, all)")
	flag.Parse()

	token := client.Authenticate()

	client.SubRedditListing(token, *username, *where, *limit, *time)
	//	if *where == "about" {
	//		client.UserAbout(token, *username)
	//	}
	//  client.UserPosts(token,*username,*where,*limit,*sort,*time)

}
