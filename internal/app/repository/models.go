package repository

import "time"

type recipe struct {
	Name      string
	Data      recipeData
	Requires  []recipeData
	CraftTime time.Duration
}

type recipeData struct {
	ID, Count uint64
}
