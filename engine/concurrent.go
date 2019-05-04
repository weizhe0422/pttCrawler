package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureWorkChan(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	var (
		in     chan Request
		out    chan ParseResult
		result ParseResult
	)

	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	in = make(chan Request)
	out = make(chan ParseResult)
	c.Scheduler.ConfigureWorkChan(in)

	for i := 0; i < c.WorkerCount; i++ {
		createWorker(in, out)
	}

	for {
		result = <-out
		for _, item := range result.Items {
			fmt.Printf("Got items: %v", item)
		}
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	var (
		request Request
		result  ParseResult
		err     error
	)
	go func() {
		for {
			request = <-in
			if result, err = worker(request); err != nil {
				continue
			}

			out <- result
		}
	}()
}
