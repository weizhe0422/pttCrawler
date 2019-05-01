package engine

import (
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/Fetcher"
	"log"
)

func Run(seeds ...Request) {
	var (
		requests    []Request
		body        []byte
		err         error
		parseResult ParseResult
	)

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		if body, err = Fetcher.Fetch(r.Url); err != nil {
			log.Printf("wrong to fetch URL %s: %v", r.Url, err)
			continue
		}

		parseResult = r.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

	}
}
