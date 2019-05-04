package Parser

import (
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/engine"
	"regexp"
)

const BoardListReg = `        <div class="b-ent">
            <a class="board" href="(/bbs/[0-9a-zA-Z]+/index.html)">
                <div class="board-name">([0-9a-zA-Z]+)</div>
                <div class="board-nuser"><span class="[^>]*">[0-9]+</span></div>
                <div class="board-class">([^<>]*)</div>
                <div class="board-title">([^<>]*)</div>
            </a>
        </div>`

const BoardUrlHead = "https://www.ptt.cc"

func ParseTopBoardList(contents []byte) engine.ParseResult {
	var (
		compile *regexp.Regexp
		matchs  [][][]byte
		result  engine.ParseResult
	)

	compile = regexp.MustCompile(BoardListReg)
	matchs = compile.FindAllSubmatch(contents, -1)

	result = engine.ParseResult{}

	for _, match := range matchs {
		result.Items = append(result.Items, "Board Name: "+string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:       BoardUrlHead + string(match[1]),
			ParseFunc: ParseArticleList, //engine.NilParseFunc,
		})
	}

	return result
}
