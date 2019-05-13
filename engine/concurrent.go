package Engine

import (
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/Fetch"
	"log"
)

type Concurrent struct {
	WorkCount int
	Scheduler Scheduler
}

type Scheduler interface {
	Submit(Request)
	WorkerReady(chan Request)
	Run()
}

func (c *Concurrent) Run(seeds ...Request) {
	var (
		out    chan ParseResult
		result ParseResult
	)

	out = make(chan ParseResult)

	c.Scheduler.Run()

	for i := 0; i < c.WorkCount; i++ {
		createWorker(out, c.Scheduler)
	}

	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	for {
		result = <-out
		for _, r := range result.Requests {
			c.Scheduler.Submit(r)
		}
		for _, item := range result.Items {
			log.Printf("Got item: %v", item)
		}
	}

}

func createWorker(out chan ParseResult, scheduler Scheduler) {
	var (
		in      chan Request
		request Request
		result  ParseResult
	)

	in = make(chan Request)

	go func() {
		for {
			scheduler.WorkerReady(in)

			request = <-in
			result = worker(request)
			out <- result
		}
	}()
}

func worker(r Request) ParseResult {
	var (
		content []byte
		err     error
	)

	if content, err = Fetch.Fetch(r.URL); err != nil {
		return ParseResult{}
	}

	return r.ParseFunc(content)
}
