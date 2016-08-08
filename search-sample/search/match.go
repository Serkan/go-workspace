package search

import (
	"fmt"
	"log"
)

type Result struct {
	Field   string
	Content string
}

type Mather interface {
	Search(feed *Feed, term string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, term string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, term)
	if err != nil {
		log.Println(
	}

	for _, result := range searchResults {
		results <- result
	}
}

func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}

func Register(feedType string, matcher Macther) {
	_, exist := matchers[feedType]
	if exist {
		log.Fatalln("There is already a matcher for this type %s", feedType)
	}

	matchers[feedType] = mather
}
