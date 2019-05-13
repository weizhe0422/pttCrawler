package Engine

type Request struct {
	URL         string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParseFunc([]byte) ParseResult {
	return ParseResult{}
}
