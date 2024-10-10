package api

import (
	"database/sql"
	"games-pet-project/pkg/api/initializers"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	components := initializers.InitializeComponents(db)

	gamesGroup := r.Group("/games")
	{
		gamesGroup.POST("/", components.GameHandler.CreateGame)
	}

	return r
}
