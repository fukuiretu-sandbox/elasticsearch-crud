package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type hoge struct {
	one   string
	two   string
	three string
}

func newHoge(record []string) *hoge {
	result := new(hoge)
	result.one = record[0]
	result.two = record[1]
	result.three = record[2]

	return result
}

func main() {
	fp, err := os.Open("fixtures/test.csv")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.Comma = ','
	reader.LazyQuotes = true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		// hoge := &hoge{record[0], record[1], record[2]}
		hoge := newHoge(record)
		fmt.Println(hoge)
	}
}
