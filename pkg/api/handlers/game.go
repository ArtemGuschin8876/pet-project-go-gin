package handlers

import (
	"games-pet-project/pkg/api/repositories"
	"games-pet-project/pkg/models"
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
	var newGame models.Game

	if err := c.BindJSON(&newGame); err != nil {
		g.GameRepository.Cfg.Logger.Err(err).Msg("error binding newGame")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	if err := g.GameRepository.Insert(&newGame); err != nil {
		g.GameRepository.Cfg.Logger.Err(err).Msg("error inserting newGame to database")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to insert game",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "game created successfully", "game": newGame})
}

func (g *GameHandler) GetGameList(c *gin.Context) {
	games, err := g.GameRepository.GetList()
	if err != nil {
		g.GameRepository.Cfg.Logger.Err(err).Msg("error getting games list")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve games",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": games,
	})
}

func (g *GameHandler) GetGameByID(c *gin.Context) {
	idParam := c.Param("id")

	game, err := g.GameRepository.GetByID(idParam)
	if err != nil {
		g.GameRepository.Cfg.Logger.Err(err).Msg("error fetching game")
		c.JSON(http.StatusNotFound, gin.H{"error": "game not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": game})
}

func (g *GameHandler) DeleteGame(c *gin.Context) {
	idParam := c.Param("id")

	err := g.GameRepository.Delete(idParam)
	if err != nil {
		g.GameRepository.Cfg.Logger.Err(err).Msg("error deleting game")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "delete game ID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "game created deleted"})

}

func (g *GameHandler) UpdateGame(c *gin.Context) {
	var updGame models.Game

	idParam := c.Param("id")

	if err := c.BindJSON(&updGame); err != nil {
		g.GameRepository.Cfg.Logger.Err(err).Msg("error binding updateGame")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	updGame.ID = idParam

	if err := g.GameRepository.Update(&updGame); err != nil {
		g.GameRepository.Cfg.Logger.Err(err).Msg("error updating updateGame")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "update game",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "game updated successfully", "game": updGame})
}
