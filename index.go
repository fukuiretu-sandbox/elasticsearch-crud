package main

import (
	"fmt"
	"gopkg.in/olivere/elastic.v3"
)

type Tweet struct {
	User     string
	Message  string
}

func main() {
	fmt.Println("start")

	client, err := elastic.NewClient(elastic.SetURL("http://192.168.33.12:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	// _, err = client.CreateIndex("twitter").Do()
	// if err != nil {
	//   // Handle error
	//   panic(err)
	// }

	// Add a document to the index
	tweet := `{"user" : "ror", "message" : "fuck you"}`
	_, err = client.Index().
		Index("twitter").
		Type("tweet").
		BodyJson(tweet).
		Do()
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("end")
}
