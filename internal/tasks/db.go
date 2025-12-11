package tasks

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB(connString string) (*pgxpool.Pool, error) {
	var err error
	DB, err = pgxpool.New(context.Background(), connString)
	return DB, err
}
