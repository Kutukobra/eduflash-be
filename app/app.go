package app

import (
	"github.com/Kutukobra/eduflash-be/app/config"
	"github.com/Kutukobra/eduflash-be/app/database"
)

type App struct {
}

func New(cfg *config.Config) (*App, error) {
	PGDatabase, err := database.NewPostgresDatabase(cfg.PGString)
	if err != nil {
		return nil, err
	}

	return &App{}, nil
}
