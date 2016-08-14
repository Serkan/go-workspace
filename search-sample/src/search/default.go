package search

import (
	"log"
)

type defaultMather struct{}

func init() {
	log.Println("Default matcher init called")
	var matcher defaultMather
	Register("default", matcher)
}

func (m defaultMather) Search(feed *Feed, term string) ([]*Result, error) {
	return nil, nil
}

