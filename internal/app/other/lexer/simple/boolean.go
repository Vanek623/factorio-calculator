package simple

import (
	"factorio-calculator/internal/app/other/models"
	"factorio-calculator/internal/app/other/parser"
)

type Boolean struct {
	it *parser.Iterator
}

func NewBoolean(it *parser.Iterator) *Boolean {
	return &Boolean{it}
}

func (b *Boolean) Parse() (bool, error) {
	if !l.HasTokens([]models.Type{models.Field, models.Assign, models.Bool}) {
		return false, errors.incorrectHandlerError
	}

	str := l.CurToken().value
	val := false
	if l.NextTokenWithSkip(1).value == "true" {
		val = true
	}

	return str, val, nil
}
