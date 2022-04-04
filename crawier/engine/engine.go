package engine

import (
	"Study/crawier/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetcher %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error, fetcher url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Request...)

		for _, item := range parseResult.Items {
			log.Printf("Get item %v", item)
		}
	}

}
