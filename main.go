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

	search := flag.String("type", "user", "one of(user,subreddit)")
	username := flag.String("user", "username", "Username you want search")
	subreddit := flag.String("subreddit", "golang", "Subreddit you want search")
	where := flag.String("where", "about", "one of (upvoted,downvoted,overview,submitted)")
	limit := flag.String("limit", "25", "the maximum number of items desired (default: 25, maximum: 100)")
	sort := flag.String("sort", "hot", "one of (hot, new, top, controversial)")
	time := flag.String("time", "week", "one of (hour, day, week, month, year, all)")
	flag.Parse()

	token := client.Authenticate()

	if *search == "user" {
		if *where == "about" {
			client.UserAbout(token, *username)
		}

		client.UserPosts(token, *username, *where, *limit, *sort, *time)
	}
	if *search == "subreddit" {
		client.SubRedditListing(token, *subreddit, *where, *limit, *time)
	}

}
