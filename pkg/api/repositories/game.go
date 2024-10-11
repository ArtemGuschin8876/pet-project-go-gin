package repositories

import (
	"database/sql"
	"games-pet-project/pkg/config"
	"games-pet-project/pkg/models"
)

type GameRepository struct {
	DB  *sql.DB
	Cfg config.Config
}

func NewGameRepository(db *sql.DB, cfg config.Config) *GameRepository {
	return &GameRepository{DB: db, Cfg: cfg}
}

func (g *GameRepository) Insert(game *models.Game) error {

	query := `
	INSERT INTO gameslist (name, article)
	VALUES ($1, $2)
	RETURNING id
	`
	args := []interface{}{
		game.Name,
		game.Article,
	}

	err := g.DB.QueryRow(query, args...).Scan(&game.ID)
	if err != nil {
		g.Cfg.Logger.Err(err).Msgf("Error executing query: %s", err)
	}

	g.Cfg.Logger.Info().Msg("message has been successfully added to the database")
	return nil
}

func (g *GameRepository) GetList() ([]models.Game, error) {

	query := `
	SELECT id, name, article FROM gameslist
	`

	rows, err := g.DB.Query(query)
	if err != nil {
		g.Cfg.Logger.Err(err).Msgf("error executing query: %v", err)
	}
	defer rows.Close()

	var games []models.Game

	for rows.Next() {
		var gm models.Game

		if err := rows.Scan(&gm.ID, &gm.Name, &gm.Article); err != nil {
			g.Cfg.Logger.Err(err).Msgf("error scanning row: %v", err)
		}

		games = append(games, gm)
	}

	return games, nil
}

func (g *GameRepository) GetByID(id string) (*models.Game, error) {
	var game models.Game

	query := `SELECT id, name, article FROM gameslist WHERE id = $1`

	err := g.DB.QueryRow(query, id).Scan(&game.ID, &game.Name, &game.Article)
	if err != nil {
		if err == sql.ErrNoRows {
			g.Cfg.Logger.Info().Msgf("No game found with ID: %s", id)
			return nil, err
		}
		g.Cfg.Logger.Err(err).Msgf("Error fetching game with ID: %s", id)
		return nil, err
	}

	return &game, nil
}

func (g *GameRepository) Update(game *models.Game) error {

	query := `
	UPDATE gameslist SET name = $1, article = $2 WHERE id = $3
	`

	_, err := g.DB.Exec(query, game.Name, game.Article, game.ID)
	if err != nil {
		g.Cfg.Logger.Err(err).Msgf("error updating game: %v", err)
	}

	return nil
}

func (g *GameRepository) Delete(id string) error {

	query := `
	DELETE FROM gameslist WHERE id = $1
	`

	res, err := g.DB.Exec(query, id)
	if err != nil {
		g.Cfg.Logger.Err(err).Msgf("error deleting message: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		g.Cfg.Logger.Err(err).Msgf("error getting affected rows: %v", err)
	}

	if rowsAffected == 0 {
		g.Cfg.Logger.Err(err).Msgf("no message found with id: %s", id)
	}

	return nil
}
