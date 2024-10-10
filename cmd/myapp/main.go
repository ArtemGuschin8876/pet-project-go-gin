package main

import (
	"context"
	"games-pet-project/pkg/config"
	postgresdb "games-pet-project/pkg/database/postgreSQL"
	"log"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env file")
	}

	srvPort := os.Getenv("PORT")
	dsnPostgres := os.Getenv("DSN_POSTGRES")

	cfg := config.NewConfig(srvPort, dsnPostgres)

	if err := realMain(cfg); err != nil {
		cfg.Logger.Err(err).Msg("error in realMain func")
	}

}

func realMain(cfg config.Config) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	db, err := postgresdb.ConnectToDB(cfg.DSN, cfg)
	if err != nil {
		cfg.Logger.Err(err).Msg("error connect to db")
	}

	defer db.Close()
	cfg.Logger.Info().Msg("database connection established")

	if err := Run(ctx, cfg, db); err != nil {
		cfg.Logger.Err(err).Msg("error running server")
		return err
	}

	return nil
}
