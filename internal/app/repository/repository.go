package repository

type repository interface {
	GetRecipe(ID uint64) (*Recipe, error)
	AddRecipe(recipe *Recipe) (uint64, error)
}
