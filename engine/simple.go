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
		request     []Request
		parseResult ParseResult
		err         error
	)
	for _, r := range seeds {
		request = append(request, r)
	}

	for len(request) > 0 {
		r := request[0]
		request = request[1:]

		if parseResult, err = worker(r); err != nil {
			continue
		}

		request = append(request, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("get item: %s", item)
		}

	}
}

func worker(r Request) (ParseResult, error) {
	var (
		content []byte
		ok      bool
	)

	log.Printf("fetching: %s", r.URL)
	if content, ok = Fetcher.Fetch(r.URL); !ok {
		log.Printf("failed to fetch %s", r.URL)
		return ParseResult{}, fmt.Errorf("failed to fetch %s", r.URL)
	}

	return r.ParseFunc(content), nil
}
