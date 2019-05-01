package main

import (
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/PTT/parser"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/engine"
)

const PTTUrl = "https://www.ptt.cc/bbs/hotboards.html"

func main() {
	engine.Run(engine.Request{
		Url:       PTTUrl,
		ParseFunc: parser.ParsetopBoardsList})
}
