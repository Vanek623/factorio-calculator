package other

import (
	"factorio-calculator/internal/app/other/parser"

	"github.com/pkg/errors"
)

type ParsedData struct {
	Object         *Token
	SimpleFields    map[Token]*Token
	DifficultFields []*ParsedData
}

type MapParser struct {
	TokenIterator
}

func (p MapParser) Parse(str string) (root *ParsedData, err error) {
	p.tokens = parser.WordParser{}.Parse(str)

	if !p.hasTokens([]TokenType{BeginData, RollO, CurvedO, CurvedO}) {
		return nil, errors.New("incorrect file begin")
	}

	root.Object = p.curToken().Clone()
	p.nextTokenWithSkip(2)
	root.DifficultFields = p.extractData(nil).DifficultFields

	return
}

func (p *MapParser) extractData(object *Token) (data *ParsedData) {
	isEnd := func(t *Token) bool {
		return t == nil || t.Type == CurvedC
	}

	data = &ParsedData{ Object: object }
	for token := p.curToken(); !isEnd(token); token = p.nextToken() {
		if f, v, err := p.extractSimpleField(); err == nil {
			data.SimpleFields[*f] = v
		} else if token.Type == CurvedO {
			p.nextToken()
			data.DifficultFields = append(data.DifficultFields, p.extractData(nil))
		} else if token.Type == Field  {
			p.nextToken()
			data.DifficultFields = append(data.DifficultFields, p.extractData(token))
		} else {
			return nil
		}
	}

	return
}

func (p *MapParser) extractSimpleField() (field, value *Token, err error) {
	if p.hasTokens([]TokenType{Field, Assign, CurvedO}) {
		return nil, nil, incorrectHandlerError
	}

	if p.hasTokens()

	return p.curToken().Clone(), p.nextTokenWithSkip(1).Clone(), nil
}

func (p *MapParser) extractDifficultField() (field *Token, data *ParsedData, err error) {



}
