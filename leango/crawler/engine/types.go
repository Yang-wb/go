package engine

type Request struct {
	Url    string
	Parser Parser
}

type ParserFunc func(contents []byte, url string) ParseResult

type Parser interface {
	Parser(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type NilParser struct {
}

func (n NilParser) Parser(contents []byte, url string) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parser(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
