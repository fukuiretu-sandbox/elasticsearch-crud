package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"

	"gopkg.in/olivere/elastic.v3"
)

type Tweet struct {
	User    string
	Message string
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s:
   %s [OPTIONS] ARGS...
Options\n`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	var word = flag.String("word", "hoge", "search word option")
	flag.Parse()

	fmt.Printf("search word is %d \n", word)

	fmt.Println("search start...")
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.33.12:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	// Search with a term query
	// 完全一致
	// termQuery := elastic.NewTermQuery("user", "retu")
	// 前方一致
	matchQuery := elastic.NewMatchPhrasePrefixQuery("user", *word)
	searchResult, err := client.Search().
		Index("twitter").   // search in index "twitter"
		Query(matchQuery).  // specify the query
		Sort("user", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do()                // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(Tweet); ok {
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	}

	fmt.Println("end")
}

// func esClient(endpoint string) Client {
//   client, err := elastic.NewClient(elastic.SetURL(endpoint))
//   if err != nil {
// 		// Handle error
// 		panic(err)
// 	}
//   return client
// }
