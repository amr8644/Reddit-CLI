package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reddit/cli/types"
)


func (client *BaseClient) SubRedditListing(token Token, subr string, where string, limit string, time string) {
	var subreddit types.SubRedditListings 
	url := "r/" + subr + "/" + where + "?limit=" + limit  + "&t=" + time

	header := []string{
		"ID", "Upvotes", "isOver18", "SubReddit", "Name", "Author", "Title", "URL",
	}

	data, err := client.Requests(token, url)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal([]byte(data), &subreddit)

	f, err := os.OpenFile("output/"+subr+"_"+where+".csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	w := csv.NewWriter(f)
	w.Write(header)

    fmt.Println(subreddit.Data.Children[0].Data.Subreddit)
}
