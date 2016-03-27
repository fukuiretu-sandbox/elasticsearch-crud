package main

import (
	"encoding/csv"
	"io"
	"os"

	"gopkg.in/olivere/elastic.v3"
)

type restaurant struct {
	id               string
	name             string
	property         string
	alphabet         string
	nameKana         string
	prefID           string
	areaID           string
	stationID1       string
	stationTime1     string
	stationDistance1 string
	stationID2       string
	stationTime2     string
	stationDistance2 string
	stationID3       string
	stationTime3     string
	stationDistance3 string
	categoryID1      string
	categoryID2      string
	categoryID3      string
	categoryID4      string
	categoryID5      string
	zip              string
	address          string
	northLatitude    string
	eastLongitude    string
	description      string
	purpose          string
	openMorning      string
	openLunch        string
	openLate         string
	photoCount       string
	specialCount     string
	menuCount        string
	fanCount         string
	accessCount      string
	createdOn        string
	modifiedOn       string
	closed           string
}

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

func newRestaurant(record []string) *restaurant {
	result := new(restaurant)
	result.id = record[0]
	result.name = record[1]
	result.property = record[2]
	result.alphabet = record[3]
	result.nameKana = record[4]
	result.prefID = record[5]
	result.areaID = record[6]
	result.stationID1 = record[7]
	result.stationTime1 = record[8]
	result.stationDistance1 = record[9]
	result.stationID2 = record[10]
	result.stationTime2 = record[11]
	result.stationDistance2 = record[12]
	result.stationID3 = record[13]
	result.stationTime3 = record[14]
	result.stationDistance3 = record[15]
	result.categoryID1 = record[16]
	result.categoryID2 = record[17]
	result.categoryID3 = record[18]
	result.categoryID4 = record[19]
	result.categoryID5 = record[20]
	result.zip = record[21]
	result.address = record[22]
	result.northLatitude = record[23]
	result.eastLongitude = record[24]
	result.description = record[25]
	result.purpose = record[26]
	result.openMorning = record[27]
	result.openLunch = record[28]
	result.openLate = record[29]
	result.photoCount = record[30]
	result.specialCount = record[31]
	result.menuCount = record[32]
	result.fanCount = record[33]
	result.accessCount = record[34]
	result.createdOn = record[35]
	result.modifiedOn = record[36]
	result.closed = record[37]
	return result
}

func toMapingData(restaurant *restaurant) *mappingData {
	result := new(mappingData)
	result.RestaurantID = restaurant.id
	result.Name = restaurant.name
	result.NameAlphabet = restaurant.alphabet
	result.NameKana = restaurant.nameKana
	result.Address = restaurant.address
	result.Description = restaurant.description
	result.Purpose = restaurant.purpose
	result.Category = restaurant.categoryID1
	result.PhotoCount = restaurant.photoCount
	result.MenuCount = restaurant.menuCount
	result.AccessCount = restaurant.accessCount
	result.Closed = restaurant.closed
	return result
}

func main() {
	fp, err := os.Open("fixtures/restaurants.csv")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.Comma = ','
	reader.LazyQuotes = true

	client, err := elastic.NewClient(elastic.SetURL("http://192.168.33.12:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}
	bulkRequest := client.Bulk()

	i := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		restaurant := newRestaurant(record)
		mapingData := toMapingData(restaurant)
		indexReq := elastic.NewBulkIndexRequest().Index("ldgourmet").Type("restaurant").Doc(mapingData)
		bulkRequest = bulkRequest.Add(indexReq)

		if i%1000 == 0 {
			bulkResponse, err := bulkRequest.Do()
			if err != nil {
				panic(err)
			}
			if bulkResponse == nil {
				panic("expected bulkResponse to be != nil; got nil")
			}
		}
		i++
	}

	if bulkRequest.NumberOfActions() > 0 {
		bulkRequest.Do()
	}
}
