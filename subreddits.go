package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"reddit/cli/types"
	"strconv"
)

func (client *BaseClient) SubRedditListing(token Token, subr string, where string, limit string, time string) {
	var subreddit types.SubRedditListings
	url := "r/" + subr + "/" + where + "?limit=" + limit + "&t=" + time

	header := []string{
		"ID", "Author", "Title", "DownVotes", "Upvotes", "Over18", "Subscribers", "URL",
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
	for _, v := range subreddit.Data.Children {
		records := []string{
			v.Data.ID,
			v.Data.Author,
			v.Data.Title,
			strconv.Itoa(v.Data.Downs),
			strconv.Itoa(v.Data.Ups),
			strconv.FormatBool(v.Data.Over18),
			strconv.Itoa(v.Data.SubredditSubscribers),
			"https://www.reddit.com/" + v.Data.Permalink,
		}
		w.Write(records)
	}
	w.Flush()

}
