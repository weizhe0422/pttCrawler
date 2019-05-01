package parser

import (
	"fmt"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/Fetcher"
	"testing"
)

func TestParsetopBoardsList(t *testing.T) {
	var (
		contents []byte
		err      error
	)
	if contents, err = Fetcher.Fetch("https://www.ptt.cc/bbs/hotboards.html"); err != nil {
		panic(err)
	}
	fmt.Println(string(contents))
}
