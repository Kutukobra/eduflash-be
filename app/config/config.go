package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PGString string
	AppPort  string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := Config{}

	config.PGString = os.Getenv("PG_STRING")
	if config.PGString == "" {
		return nil, errors.New("no pg connection string found")
	}

	config.AppPort = os.Getenv("APP_PORT")
	if config.AppPort == "" {
		config.AppPort = "3000"
	}

	return &config, nil
}
