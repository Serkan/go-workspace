package search

import (
	"encoding/json"
	"os"
)

const dataFile = "/home/serkan/dev/go-workspace/search-sample/src/data/data.json"

type Feed struct {
	Name string `json:site`
	URI  string `json:link`
	Type string `json:type`
}

func GetFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
