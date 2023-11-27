package parser

type Tuple struct {
	min int
	max int
}

func NewTuple(min int, max int) *Tuple {
	return &Tuple{
		min,
		max,
	}
}

type Players struct {
	Name       string `json:"name"`
	High_score int    `json:"high_score"`
}

type ParserInterface interface {
	ParseFile(file string) ([]Players, error)
}

type Parser struct {
	Parser ParserInterface
}

func (p *Parser) ParseFile(fileName string) ([]Players, error) {
	return p.Parser.ParseFile(fileName)
}

func NewParser(p ParserInterface) *Parser {
	return &Parser{
		Parser: p,
	}
}
