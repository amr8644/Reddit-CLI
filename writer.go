package main

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteToCSV(header []string, records []string) {

	f, err := os.OpenFile("output/data.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalln(err)
	}

	w := csv.NewWriter(f)
	w.Write(header)

	// for _, v := range user.Data.Children {
	// 	records := []string{
	// 		v.data.id, v.data.subreddit, v.data.name, v.data.author, v.data.title, v.data.permalink,
	// 	}
	// 	w.Write(records)
	// }
}
