package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	//canal sem buffer
	results := make(chan *Result)

	//grupo de espera
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	// goroutine para cada feed
	for _, feed := range feeds {
		// obtem o matcher
		matcher, exits := matchers[feed.Type]
		if !exits {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
		}(matcher, feed)
	}

	// goroutine de espera
	go func() {
		waitGroup.Wait()
		close(results)
	}()

	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
