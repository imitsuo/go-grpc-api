package databases

import (
	"context"

	"go-grpc-api/internal/configs"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func ConnectDB() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(configs.GetSettings()["POSTGRES"])
	if err != nil {
		panic(err)
	}
	config.MaxConns = 150
	config.MinConns = 100

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}
	return conn

}

func GetDB() *pgxpool.Pool {
	if db == nil {
		db = ConnectDB()
	}

	return db
}
