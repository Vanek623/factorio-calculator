package models

type Recipe struct {
	Type, Name, Category string
	DP                   map[DifficultValue]*DifficultParams
}

type DifficultValue int

const (
	Normal DifficultValue = iota
	Expensive
)

type DifficultParams struct {
	Enabled     bool
	Ingredients []Ingredient
	Time        float64
	Result      string
	Count       int64
}

type Ingredient struct {
	Type   string
	Name   string
	Amount int64
}
