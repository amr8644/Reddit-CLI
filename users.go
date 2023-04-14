package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reddit/cli/types"
	"strconv"
)

func (client *BaseClient) UserAbout(token Token, username string) {

	var user types.UserAbout
	url := "user/" + username + "/about"
	header := []string{
		"ID", "Name", "Karma", "URL",
	}

	data, err := client.Requests(token, url)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal([]byte(data), &user)

	f, err := os.OpenFile("output/"+username+".csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	w := csv.NewWriter(f)
	w.Write(header)

	records := []string{
		user.Data.ID,user.Data.Name ,strconv.Itoa(user.Data.TotalKarma),
	}

	w.Write(records)
    fmt.Println(records)
}

func (client *BaseClient) UserPosts(token Token, username string, where string, limit string, sort string, t string) {

	var user types.UserUpvotes
	url := "user/" + username + "/" + where + "?limit=" + limit + "&sort=" + sort + "&t=" + t

	header := []string{
		"ID", "Upvotes", "isOver18", "SubReddit", "Name", "Author", "Title", "URL",
	}

	data, err := client.Requests(token, url)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal([]byte(data), &user)

	f, err := os.OpenFile("output/"+username+"_"+where+".csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	w := csv.NewWriter(f)
	w.Write(header)

	for _, v := range user.Data.Children {
		records := []string{
			v.Data.ID, strconv.Itoa(v.Data.Ups), strconv.FormatBool(v.Data.Over18), v.Data.Subreddit, v.Data.Author, v.Data.Title, "https://www.reddit.com/" + v.Data.Permalink,
		}
		w.Write(records)
	}
}
