package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"

	"gopkg.in/olivere/elastic.v3"
)

type mappingData struct {
	RestaurantID string `json:"restaurant_id"`
	Name         string `json:"name"`
	NameAlphabet string `json:"name_alphabet"`
	NameKana     string `json:"name_kana"`
	Address      string `json:"address"`
	Description  string `json:"description"`
	Purpose      string `json:"purpose"`
	Category     string `json:"category"`
	PhotoCount   string `json:"photo_count"`
	MenuCount    string `json:"menu_count"`
	AccessCount  string `json:"access_count"`
	Closed       string `json:"closed"`
	Location     string `json:"location"`
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
  Usage of %s:
   %s [OPTIONS] ARGS...
  Options\n`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	word := flag.String("word", "hoge", "search word option")
	limit := flag.Int("limit", 100, "search limit option")
	flag.Parse()
	fmt.Printf("search word is %d \n", *word)

	fmt.Println("search start...")
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.33.12:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	q := elastic.NewQueryStringQuery(*word)
	q = q.DefaultField("name")
	searchResult, err := client.Search().
		Index("ldgourmet").   // search in index "twitter"
		Query(q).             // specify the query
		Sort("name", true).   // sort by "user" field, ascending
		From(0).Size(*limit). // take documents 0-9
		Pretty(true).         // pretty print request and response JSON
		Do()                  // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	var ttyp mappingData
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(mappingData); ok {
			fmt.Printf("mappingData by %d", t.Name)
		}
	}
}
