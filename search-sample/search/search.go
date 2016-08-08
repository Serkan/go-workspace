package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Mather)

func Run(term string) {
	feeds, err := GetFeeds()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		mather, exist := matchers[feed.Type]
		if !exist {
			mather := matchers["default"]
		}
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, term, result)
			waitGroup.Done()
		}(matcher, feed)
	}

	waitGroup.Wait()
	close(results)

	Display(results)
}
