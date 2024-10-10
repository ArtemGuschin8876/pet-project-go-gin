package initializers

import (
	"database/sql"
	"games-pet-project/pkg/api/handlers"
	"games-pet-project/pkg/api/repositories"
)

type AppComponents struct {
	GameHandler *handlers.GameHandler
}

// Инициализируем компоненты сущностей
func InitializeComponents(db *sql.DB) *AppComponents {

	gameRep := repositories.NewGameRepository(db)
	gameHandler := handlers.NewGameHandler(gameRep)

	return &AppComponents{
		GameHandler: gameHandler,
	}
}
