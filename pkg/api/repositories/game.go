package repositories

import (
	"database/sql"
	"games-pet-project/pkg/models"
)

type GameRepository struct {
	DB *sql.DB
}

func NewGameRepository(db *sql.DB) *GameRepository {
	return &GameRepository{DB: db}
}

func (g *GameRepository) Add(game *models.Game) error {
	return nil
}
