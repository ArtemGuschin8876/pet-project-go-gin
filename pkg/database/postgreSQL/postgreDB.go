package postgresdb

import (
	"context"
	"database/sql"
	"games-pet-project/pkg/config"
	"time"

	_ "github.com/lib/pq"
)

const driverPostgres = "postgres"

func ConnectToDB(dsn string, cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open(driverPostgres, dsn)
	if err != nil {
		cfg.Logger.Err(err).Msg("error openning database")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		cfg.Logger.Err(err).Msg("error ping database")
	}

	return db, nil
}
