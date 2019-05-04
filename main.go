package main

import (
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/PTT/Parser"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/engine"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/scheduler"
)

const PTTHotBoard = "https://www.ptt.cc/bbs/hotboards.html"

func main() {
	crawlerEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}

	//crawlerEngine := engine.SimpleEngine{}

	crawlerEngine.Run(engine.Request{
		URL:       PTTHotBoard,
		ParseFunc: Parser.ParseTopBoardList,
	})
}
