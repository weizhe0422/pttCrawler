package scheduler

import "github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureWorkChan(c chan engine.Request) {
	s.workChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func(){
		s.workChan <- r
	}()
}