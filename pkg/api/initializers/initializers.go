package initializers

import (
	"database/sql"
	"games-pet-project/pkg/api/handlers"
	"games-pet-project/pkg/api/repositories"
	"games-pet-project/pkg/config"
)

type AppComponents struct {
	GameHandler *handlers.GameHandler
}

// Инициализируем компоненты сущностей
func InitializeComponents(db *sql.DB, cfg config.Config) *AppComponents {

	gameRep := repositories.NewGameRepository(db, cfg)
	gameHandler := handlers.NewGameHandler(gameRep)

	return &AppComponents{
		GameHandler: gameHandler,
	}
}
