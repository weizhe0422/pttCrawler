package engine

import "log"

type Concurrent struct {
	Scheduler Scheduler
	WorkCount int
}


type Scheduler interface {
	Submit(Request)
	ConfigureWorkChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (c *Concurrent) Run (seeds ...Request){
	var(
		out chan ParseResult
		result ParseResult
	)

	out = make(chan ParseResult)
	c.Scheduler.Run()

	for i := 0; i < c.WorkCount; i++{
		createWorker(out, c.Scheduler)
	}

	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	for{
		result = <- out

		for _, item := range result.Items {
			log.Printf("Got item: %v", item)
		}

		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(out chan ParseResult, s Scheduler){
	var(
		request Request
		parseResult ParseResult
		err error
		in chan Request
	)
	in = make(chan Request)

	go func(){
		for{
			s.WorkerReady(in)
			request = <- in
			if parseResult, err = worker(request); err!=nil{
				continue
			}
			out <- parseResult
		}
	}()
}