package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	error2 "levi-api/model/error"
	"os"
)

func Connect() (*pgxpool.Pool, error) {
	var dbEr error2.DBError
	if pool, err := pgxpool.New(context.Background(), os.Getenv("DB_URI")); err != nil {
		return nil, dbEr.Connect(err.Error())
	} else {
		return pool, nil
	}
}
