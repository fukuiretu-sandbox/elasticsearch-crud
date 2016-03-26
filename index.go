package main

import (
	"fmt"

	"gopkg.in/olivere/elastic.v3"
)

type tweet struct {
	user    string
	message string
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
	// tweet := `{"user" : "ror", "message" : "fuck you"}`
	// single
	// _, err = client.Index().
	// 	Index("twitter").
	// 	Type("tweet").
	// 	BodyJson(tweet).
	// 	Do()
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }

	// bulk
	tweet1 := tweet{user: "olivere", message: "Welcome to Golang and Elasticsearch."}
	tweet2 := tweet{user: "sandrae", message: "Dancing all night long. Yeah."}
	index1Req := elastic.NewBulkIndexRequest().Index("twitter").Type("tweet").Doc(tweet1)
	index2Req := elastic.NewBulkIndexRequest().Index("twitter").Type("tweet").Doc(tweet2)

	bulkRequest := client.Bulk()
	bulkRequest = bulkRequest.Add(index1Req)
	bulkRequest = bulkRequest.Add(index2Req)

	// if bulkRequest.NumberOfActions() != 2 {
	// 	panic("expected bulkRequest.NumberOfActions %d; got %d", 3, bulkRequest.NumberOfActions())
	// }

	bulkResponse, err := bulkRequest.Do()
	if err != nil {
		panic(err)
	}
	if bulkResponse == nil {
		panic("expected bulkResponse to be != nil; got nil")
	}

	// if bulkRequest.NumberOfActions() != 0 {
	// 	panic("expected bulkRequest.NumberOfActions %d; got %d", 0, bulkRequest.NumberOfActions())
	// }

	fmt.Println("end")
}
