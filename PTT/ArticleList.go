package PTT

import (
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/Engine"
	"regexp"
)

const ArticleListRegexp = `<a href="(/bbs/[A-Za-z0-9]+/M.[0-9]+.A.[a-zA-Z0-9]+.html)">([^<>]*)</a>`
const ArticleURLHead = "https://www.ptt.cc"

func ParseArticleList(content []byte) Engine.ParseResult {
	var (
		compile *regexp.Regexp
		matchs  [][][]byte
		result  Engine.ParseResult
	)

	compile = regexp.MustCompile(ArticleListRegexp)
	matchs = compile.FindAllSubmatch(content, -1)

	result = Engine.ParseResult{}
	for _, match := range matchs {
		result.Items = append(result.Items, string(match[2]))
		result.Requests = append(result.Requests, Engine.Request{
			URL:       ArticleURLHead + string(match[1]),
			ParseFunc: Engine.NilParseFunc,
		})
	}

	return result
}
