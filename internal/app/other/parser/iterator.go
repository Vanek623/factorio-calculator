package parser

import (
	"factorio-calculator/internal/app/other/models"
)

type Iterator struct {
	tokens []models.Token

	offsetID int // use for offset
	offsets  []int
}

func NewIterator(tokens []models.Token) *Iterator {
	return &Iterator{
		tokens: tokens,
	}
}

func (it *Iterator) Front() *models.Token {
	it.offsetID = 0
	return it.CurToken()
}

func (it *Iterator) Back() *models.Token {
	it.offsetID = len(it.tokens) - 1
	return it.CurToken()
}

// saveOffset сохранить офсет в стек
func (it *Iterator) saveOffset() {
	it.offsets = append(it.offsets, it.offsetID)
}

func (it *Iterator) hasOffset() bool {
	return len(it.offsets) == 0
}

// loadOffset загрузить офсет из стека
func (it *Iterator) loadOffset() {
	it.offsetID = it.offsets[len(it.offsets)-1]
	it.offsets = it.offsets[:len(it.offsets)-1]
}

// dropOffset удалить последний офсет
func (it *Iterator) dropOffset() {
	it.offsets = it.offsets[:len(it.offsets)-1]
}

func (it *Iterator) CurToken() *models.Token {
	if it.offsetID < 0 || it.offsetID >= len(it.tokens) {
		return nil
	}

	return &it.tokens[it.offsetID]
}

func (it *Iterator) NextToken() *models.Token {
	return it.NextTokenWithSkip(0)
}

func (it *Iterator) PrevToken() *models.Token {
	return it.PrevTokenWithSkip(0)
}

func (it *Iterator) PrevTokenWithSkip(skipCount int) *models.Token {
	it.offsetID -= skipCount + 1

	return it.CurToken()
}

func (it *Iterator) NextTokenWithSkip(skipCount int) *models.Token {
	it.offsetID += skipCount + 1

	return it.CurToken()
}

func (it *Iterator) HasTokens(tokens []models.Type) bool {
	it.saveOffset()
	defer it.loadOffset()

	for i, token := 0, it.CurToken(); token != nil && i < len(tokens); token = it.NextToken() {
		if token.Type != tokens[i] {
			return false
		}
		i++
	}

	return true
}

func (it *Iterator) Clone() *Iterator {
	newIt := *it
	return &newIt
}
