package engine

import (
	"fmt"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/Fetcher"
	"log"
)


type SimpleEngine struct {

}

func (SimpleEngine) Run(seeds ...Request) {
	var (
		requests    []Request
		r           Request
		err         error
		parseResult ParseResult
	)

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r = requests[0]
		requests = requests[1:]

		if parseResult, err = worker(r); err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Get item: %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	var (
		content []byte
		ok      bool
	)
	if content, ok = Fetcher.Fetch(r.URL); !ok {
		return ParseResult{}, fmt.Errorf("failed to fetch: %s", r.URL)
	}

	return r.ParseFunc(content), nil
}
