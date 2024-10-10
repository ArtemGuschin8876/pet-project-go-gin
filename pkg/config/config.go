package config

import (
	"games-pet-project/pkg/middleware/logger"

	"github.com/rs/zerolog"
)

type Config struct {
	Port   string
	DSN    string
	Logger zerolog.Logger
}

func NewConfig(srvPort, dsn string) Config {
	return Config{
		Port:   srvPort,
		Logger: logger.NewLogger(),
		DSN:    dsn,
	}
}
