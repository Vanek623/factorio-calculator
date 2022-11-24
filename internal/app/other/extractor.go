package other

import (
	"fmt"
	"strconv"

	"factorio-calculator/internal/app/other/parser"

	"github.com/pkg/errors"
)

type iLexer interface {
	extractRecipe() (Recipe, error)
	extractNumIntField() (string, int64, error)
	extractNumFloatField() (string, float64, error)
	extractStringField() (string, string, error)
	extractBoolField() (string, bool, error)
	extractDifficult() (difficult, error)
	extractIngredients() ([]Ingredient, error)
}

type Lexer struct {
	iTokenIterator
}

var fieldParseError = errors.New("cannot parse field")
var incorrectHandlerError = errors.New("this handler cannot parse this field")

func NewTokenParser() Lexer {
	return Lexer{}
}

func (l Lexer) Parse(str string) ([]Recipe, error) {
	l.iTokenIterator = NewTokenIterator(parser.WordParser{}.Parse(str))

	if !l.HasTokens([]TokenType{BeginData, RollO, CurvedO, CurvedO}) {
		return nil, errors.New("incorrect begin of file")
	}
	l.NextTokenWithSkip(2)
	var out []Recipe
	for r, err := l.extractRecipe(); err == nil; r, err = l.extractRecipe() {
		if err != nil {
			return nil, err
		}

		out = append(out, r)
		l.NextToken()
	}

	for _, r := range out {
		fmt.Printf("%+v", r)
	}

	return out, nil
}

func (l *Lexer) clear() {
	l.iTokenIterator = nil
}

func (l *Lexer) extractRecipe() (Recipe, error) {
	if !l.HasTokens([]TokenType{CurvedO}) {
		return Recipe{}, errors.New("incorrect begin of recipe")
	}

	r := Recipe{}
	r.DP = make(map[DifficultValue]*DifficultParams)
	r.DP[Normal] = &DifficultParams{}

	stringFields[nameField] = &r.Name
	stringFields[typeField] = &r.Type
	stringFields[categoryField] = &r.Category

	intFields[resultField] = &r.DP[Normal].Count

	for token := l.NextToken(); token != nil && token.Type != CurvedC; token = l.NextToken() {
		if f, v, err := l.extractNumIntField(); err == nil {

		} else if err != incorrectHandlerError {
			return Recipe{}, err
		}

		if f, v, err := l.extractNumFloatField(); err == nil {

		} else if err != incorrectHandlerError {
			return Recipe{}, err
		}

		if f, v, err := l.extractStringField(); err == nil {

		} else if err != incorrectHandlerError {
			return Recipe{}, err
		}

		if f, v, err := l.extractBoolField(); err == nil {

		} else if err != incorrectHandlerError {
			return Recipe{}, err
		}

		if d, err := l.extractDifficult(); err == nil {

		} else if err != incorrectHandlerError {
			return Recipe{}, err
		}

		if ingrs, err := l.extractIngredients(); err == nil {

		} else {
			return Recipe{}, err
		}
	}

	return r, nil
}

func (l *Lexer) extractNumIntField() (field string, value int64, err error) {
	if !l.HasTokens([]TokenType{Field, Assign, NumInt}) {
		err = incorrectHandlerError
		return
	}

	field = l.CurToken().value
	value, err = strconv.ParseInt(l.NextTokenWithSkip(1).value, 10, 64)

	return
}

func (l *Lexer) extractNumFloatField() (field string, value float64, err error) {
	if !l.HasTokens([]TokenType{Field, Assign, NumFloat}) {
		err = incorrectHandlerError
		return
	}

	field = l.CurToken().value
	value, err = strconv.ParseFloat(l.NextTokenWithSkip(1).value, 64)

	return
}

func (l *Lexer) extractStringField() (string, string, error) {
	if !l.HasTokens([]TokenType{Field, Assign, String}) {
		return "", "", incorrectHandlerError
	}

	return l.CurToken().value, l.NextTokenWithSkip(1).value, nil
}

func (l *Lexer) extractBoolField() (string, bool, error) {

}

type difficult struct {
	Val    DifficultValue
	Params DifficultParams
}

func (l *Lexer) extractDifficult() (difficult, error) {
	if !l.HasTokens([]TokenType{Field, Assign, CurvedO}) {
		return difficult{}, incorrectHandlerError
	}

	out := difficult{}
	if l.CurToken().value == normalField {
		out.Val = Normal
	} else if l.CurToken().value == expensiveField {
		out.Val = Expensive
	} else {
		return difficult{}, incorrectHandlerError
	}

	for token := l.NextTokenWithSkip(2); token != nil && token.Type != CurvedC; token = l.NextToken() {
		if err := l.extractDifficultParam(&out.Params); err != nil {
			return difficult{}, err
		}
	}

	return out, nil
}

func (l *Lexer) extractDifficultParam(params *DifficultParams) error {
	if f, v, err := l.extractNumIntField(); err == nil {
		if f == timeField {
			params.Time = float64(v)
		} else if f == countField {
			params.Count = v
		}

		return nil
	} else if err != incorrectHandlerError {
		return err
	}

	if f, v, err := l.extractNumFloatField(); err == nil {
		if f == timeField {
			params.Time = v
		}

		return nil
	} else if err != incorrectHandlerError {
		return err
	}

	if f, v, err := l.extractStringField(); err == nil {
		if f == resultField {
			params.Result = v
		}

		return nil
	} else if err != incorrectHandlerError {
		return err
	}

	if f, v, err := l.extractBoolField(); err == nil {
		if f == enabledField {
			params.Enabled = v
		}

		return nil
	} else if err != incorrectHandlerError {
		return err
	}

	if items, err := l.extractIngredients(); err == nil {
		params.Ingredients = items

		return nil
	} else {
		return err
	}
}

func (l *Lexer) extractIngredients() (items []Ingredient, err error) {
	if !l.HasTokens([]TokenType{Field, Assign, CurvedO}) || l.CurToken().value != ingredientsField {
		err = incorrectHandlerError
		return
	}

	opened := 1
	for token := l.NextTokenWithSkip(2); token != nil && opened > 0; token = l.NextToken() {
		if token.Type == CurvedO {
			opened++
		} else if token.Type == CurvedC {
			opened--
		} else if item, err := l.extractIngredient(); err == nil {
			items = append(items, item)
		} else {
			return nil, err
		}
	}

	return
}

func (l *Lexer) extractIngredient() (item Ingredient, err error) {
	if item, err = l.extractIngredientSimple(); err == nil {
		return
	} else if err != incorrectHandlerError {
		return
	}

	if item, err = l.extractIngredientDetailed(); err == nil {
		return
	} else if err != incorrectHandlerError {
		return
	}

	return
}

func (l *Lexer) extractIngredientSimple() (item Ingredient, err error) {
	if !l.HasTokens([]TokenType{String, NumInt}) {
		err = incorrectHandlerError
		return
	}

	item.Type = "item"
	item.Name = l.CurToken().value
	item.Amount, err = strconv.ParseInt(l.NextToken().value, 10, 64)

	return
}

func (l *Lexer) extractIngredientDetailed() (Ingredient, error) {
	token := l.CurToken()
	item := Ingredient{}
	for ; token != nil && token.Type != CurvedC; token = l.NextToken() {
		if f, v, err := l.extractStringField(); err == nil {
			if f == nameField {
				item.Name = v
			} else if f == typeField {
				item.Type = v
			}

			continue
		} else if err != incorrectHandlerError {
			return Ingredient{}, err
		}

		if f, v, err := l.extractNumIntField(); err == nil {
			if f == "amount" {
				item.Amount = v
			}

			continue
		} else {
			return Ingredient{}, err
		}
	}

	if token == nil {
		return Ingredient{}, fieldParseError
	}

	return item, nil
}
