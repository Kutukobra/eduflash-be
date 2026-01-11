package database

import (
	"context"

	"github.com/jackc/pgx/v5"

	_ "github.com/lib/pq"
)

func NewPostgresDatabase(connectionString string) (*pgx.Conn, error) {
	driver, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	return driver, nil
}
