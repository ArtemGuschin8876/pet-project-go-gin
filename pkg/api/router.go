package api

import (
	"database/sql"
	"games-pet-project/pkg/api/initializers"
	"games-pet-project/pkg/config"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB, cfg config.Config) *gin.Engine {
	r := gin.Default()

	components := initializers.InitializeComponents(db, cfg)

	gamesGroup := r.Group("/games")
	{
		gamesGroup.POST("/create", components.GameHandler.CreateGame)
		gamesGroup.GET("/list", components.GameHandler.GetGameList)
		gamesGroup.GET("/:id", components.GameHandler.GetGameByID)
		gamesGroup.DELETE("/:id", components.GameHandler.DeleteGame)
		gamesGroup.PUT("/:id", components.GameHandler.UpdateGame)
	}

	return r
}
