package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func New(dsn string) error {
	config, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		return fmt.Errorf("erro ao fazer o parse da configuração: %w", err)
	}

	config.MaxConns = 255
	config.MinConns = 5
	config.MaxConnLifetime = 10 * time.Minute
	config.MaxConnIdleTime = 5 * time.Minute
	config.HealthCheckPeriod = 30 * time.Second

	pool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		return fmt.Errorf("erro ao criar o pool de conexões: %w", err)
	}

	db = pool
	return nil
}

func GetConn() *pgxpool.Conn {
	conn, err := db.Acquire(context.Background())

	if err != nil {
		log.Fatalf("erro ao criar conexão com o banco de dados", err.Error())
	}

	return conn
}
