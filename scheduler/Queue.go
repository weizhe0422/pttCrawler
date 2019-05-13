package Scheduler

import (
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/Engine"
	"log"
)

type Scheduler struct {
	RequestChan chan Engine.Request
	WorkerChan  chan chan Engine.Request
}

func (s *Scheduler) Submit(r Engine.Request) {
	s.RequestChan <- r
}

func (s *Scheduler) WorkerReady(w chan Engine.Request) {
	s.WorkerChan <- w
}

func (s *Scheduler) Run() {
	log.Printf("Schedule Run")
	s.RequestChan = make(chan Engine.Request)
	s.WorkerChan = make(chan chan Engine.Request)
	go func() {
		var (
			RequestQ      []Engine.Request
			WorkerQ       []chan Engine.Request
		)

		for {
			var(
				ActiveRequest Engine.Request
				ActiveWorker  chan Engine.Request
			)

			if len(RequestQ) > 0 && len(WorkerQ) > 0 {
				ActiveRequest = RequestQ[0]
				ActiveWorker = WorkerQ[0]
			}

			select {
			case r := <-s.RequestChan:
				RequestQ = append(RequestQ, r)
			case w := <-s.WorkerChan:
				WorkerQ = append(WorkerQ, w)
			case ActiveWorker <- ActiveRequest:
				RequestQ = RequestQ[1:]
				WorkerQ = WorkerQ[1:]
			}
		}
	}()
}
