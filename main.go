package main

import (
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/PTT/parser"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/engine"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/scheduler"
)

const PTTHotBoard = "https://www.ptt.cc/bbs/hotboards.html"

func main() {
	/*engine.SimpleEngine{}.Run(engine.Request{
		URL:       PTTHotBoard,
		ParseFunc: parser.ParseHotBoardList,
	})*/

	crawlerEngine := engine.Concurrent{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 10,
	}

	crawlerEngine.Run(engine.Request{
		URL:       PTTHotBoard,
		ParseFunc: parser.ParseHotBoardList,
	})
}
