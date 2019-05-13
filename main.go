package main

import (
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/Engine"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/PTT"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/Scheduler"
)

const PTTHotBoard = "https://www.ptt.cc/bbs/hotboards.html"

func main() {

	PTTCrawler := Engine.Concurrent{
		WorkCount: 10,
		Scheduler: &Scheduler.Scheduler{},
	}

	PTTCrawler.Run(Engine.Request{
		URL:       PTTHotBoard,
		ParseFunc: PTT.ParseHotBoardList,
	})
}
