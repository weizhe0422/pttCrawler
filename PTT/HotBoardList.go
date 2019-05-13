package PTT

import (
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/Engine"
	"log"
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

func ParseHotBoardList (content []byte) Engine.ParseResult{
	var(
		compiler *regexp.Regexp
		matchs [][][]byte
		result Engine.ParseResult
	)

	compiler = regexp.MustCompile(BoardListReg)
	matchs = compiler.FindAllSubmatch(content, -1)

	result = Engine.ParseResult{}
	for _, match := range matchs{
		log.Printf("URL: %s ",BoardUrlHead + string(match[1]))
		result.Requests = append(result.Requests, Engine.Request{
			URL:       BoardUrlHead + string(match[1]),
			ParseFunc: ParseArticleList,
		})
		result.Items = append(result.Items, string(match[2]))
	}

	return result
}
