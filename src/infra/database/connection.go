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

func New(dsn string) (*pgxpool.Pool, error) {
	db, err = pgxpool.New(context.Background(), dsn)

	if err != nil {
		log.Fatalf("erro ao criar conexão com o banco de dados", err.Error())
	}

	return db, nil
}

func GetConn() *pgxpool.Conn {
	conn, err := db.Acquire(context.Background())

	if err != nil {
		log.Fatalf("erro ao criar conexão com o banco de dados", err.Error())
	}

	return conn
}
