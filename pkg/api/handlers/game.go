package handlers

import (
	"games-pet-project/pkg/api/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	GameRepository *repositories.GameRepository
}

func NewGameHandler(rep *repositories.GameRepository) *GameHandler {
	return &GameHandler{GameRepository: rep}
}

func (g *GameHandler) CreateGame(c *gin.Context) {
	c.String(http.StatusOK, "SDSASADD")
}
