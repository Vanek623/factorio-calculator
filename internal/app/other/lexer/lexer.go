package lexer

import (
	"factorio-calculator/internal/app/other/models"
	"factorio-calculator/internal/app/other/parser"
)

type Lexer struct {
	it *parser.Iterator
}

func NewLexer(tokens []models.Token) *Lexer {
	return &Lexer{parser.NewIterator(tokens)}
}

func (l *Lexer) Parse() {
	l.it.Front()
	if !l.it.HasTokens([]models.Type{BeginData, RollO, CurvedO, CurvedO})
}
