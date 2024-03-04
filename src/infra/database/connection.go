package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	db  *pgxpool.Pool
	err error
)

func New(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		panic(err)
	}

	config.MinConns = 5
	config.MaxConns = 25

	db, err = pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		panic(err)
	}

	return db, nil
}

func GetConn(ctx context.Context) *pgxpool.Conn {
	conn, err := db.Acquire(ctx)

	if err != nil {
		log.Fatalf("erro ao criar conexão com o banco de dados", err.Error())
	}

	return conn
}
