package postgres

import (
	"factorio-calculator/internal/app/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository interface {
	GetRecipe(ID uint64) (*models.Recipe, error)
	AddRecipe(recipe *models.Recipe) error
}

type Postgres struct {
	pool *pgxpool.Pool
}

func NewPostgres(pool *pgxpool.Pool) *Postgres {
	return &Postgres{pool: pool}
}

func (p *Postgres) GetRecipe(ID uint64) (*models.Recipe, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) AddRecipe(recipe *models.Recipe) error {
	//TODO implement me
	panic("implement me")
}
