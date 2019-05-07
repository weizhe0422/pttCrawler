package parser

import (
	"github.com/weizhe0422/GolangPracticeProject/FromMoocsAgain/crawler/engine"
	"regexp"
)

const ArticleListRegexp = `<a href="(/bbs/[A-Za-z0-9]+/M.[0-9]+.A.[a-zA-Z0-9]+.html)">([^<>]*)</a>`
const ArticleURLHead = "https://www.ptt.cc"

func ParseArticleList(content []byte) engine.ParseResult {
	var (
		compile *regexp.Regexp
		matchs  [][][]byte
		result  engine.ParseResult
	)

	compile = regexp.MustCompile(ArticleListRegexp)
	matchs = compile.FindAllSubmatch(content, -1)

	result = engine.ParseResult{}

	for _, match := range matchs {
		result.Items = append(result.Items, "Article title: "+string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:       ArticleURLHead + string(match[1]),
			ParseFunc: engine.NilParseFunc,
		})
	}

	return result
}
