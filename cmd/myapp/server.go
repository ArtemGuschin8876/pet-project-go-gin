package main

import (
	"context"
	"database/sql"
	"games-pet-project/pkg/api"
	"games-pet-project/pkg/config"
	"net/http"
)

// Инициализируется и запускается сервер
func Run(ctx context.Context, cfg config.Config, db *sql.DB) error {
	s := http.Server{
		Addr:    cfg.Port,
		Handler: api.NewRouter(db, cfg),
	}

	go func() {
		<-ctx.Done()
		cfg.Logger.Info().Msg("shutdown server")
		s.Shutdown(ctx)
	}()

	cfg.Logger.Info().Msgf("server starting on %v", cfg.Port)

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		cfg.Logger.Err(err).Msg("error listenning server")
		return err
	}
	return nil
}
