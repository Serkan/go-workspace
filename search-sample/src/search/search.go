package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(term string) {
	feeds, err := GetFeeds()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exist := matchers[feed.Type]
		if !exist {
			matcher = matchers["default"]
			log.Println("Matcher not found, default will be used")
		}else{
			log.Println("Matcher found " + feed.Name)
		}
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, term, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	Display(results)
}
