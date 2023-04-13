package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reddit/cli/types"
)

func (client *BaseClient) UserAbout(token Token, username string) types.UserAbout {

	var user types.UserAbout

	url := "user/" + username + "/about"

	data, err := client.Requests(token, url)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal([]byte(data), &user)

	return user
}

func (client *BaseClient) UserUpvotes(token Token, username string) types.UserUpvotes {

	var user types.UserUpvotes

	url := "user/" + username + "/upvoted"

	//  records := [][]string{
	//	{"ID", "SubRedditId", "Name", "Author", "Over18", "SubReddit", "Title", "URL"},
	//}

	data, err := client.Requests(token, url)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal([]byte(data), &user)

	//	f, err := os.OpenFile("data.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	for i, v := range user.Data.Children {
		fmt.Println(i,v)
	}
	//	w := csv.NewWriter(f)
	//	w.WriteAll(records)

	//q := user.FilterEssentialsData()

	return user
}
