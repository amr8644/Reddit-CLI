package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"reddit/cli/types"
)

func (client *BaseClient) UserUpvotes(token Token, username string) {

	var user types.UserUpvotes
	url := "user/" + username + "/upvoted"

	header := []string{
		"ID", "SubReddit", "Name", "Author", "SubReddit", "Title", "URL",
	}

	data, err := client.Requests(token, url)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal([]byte(data), &user)

	f, err := os.OpenFile("output/"+username+".csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	w := csv.NewWriter(f)
	w.Write(header)

	for _, v := range user.Data.Children {
		records := []string{
			v.Data.ID, v.Data.Subreddit, v.Data.Name, v.Data.Author, v.Data.Title, "https://www.reddit.com/" + v.Data.Permalink,
		}
		w.Write(records)
	}
}
