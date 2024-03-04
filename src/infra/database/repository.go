package database

import (
	"context"
	"crebito/src/domain"

	"github.com/jackc/pgx/v5"
)

type SqlRepository struct{}

func NewSqlRepositories() *SqlRepository {
	return &SqlRepository{}
}

func (r *SqlRepository) GetClient(clientID int32) (*domain.Client, error) {
	conn := GetConn()
	defer conn.Release()

	result := conn.QueryRow(
		context.Background(),
		`SELECT ACCOUNT_LIMIT, BALANCE FROM CLIENTS WHERE ID = $1;`,
		clientID,
	)

	client := domain.Client{ID: clientID}

	if err := result.Scan(
		&client.AccountLimit,
		&client.Balance,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrClientNotFound
		}
		return nil, err
	}

	return &client, nil
}

func (r *SqlRepository) FindLastTransactionsByClient(clientID int32) ([]domain.Transaction, error) {
	conn := GetConn()
	defer conn.Release()

	rows, err := conn.Query(
		context.Background(),
		`SELECT VALUE, TYPE, DESCRIPTION, CREATED_AT 
		FROM TRANSACTIONS 
		WHERE CLIENT_ID = $1 
		ORDER BY ID DESC LIMIT 10;`,
		clientID,
	)

	if err != nil {
		return nil, err
	}

	var result []domain.Transaction

	for rows.Next() {
		var t domain.Transaction
		if err := rows.Scan(
			&t.Value,
			&t.Type,
			&t.Description,
			&t.CreatedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (r *SqlRepository) CreateTransactionAndUpdateBalance(client *domain.Client, t *domain.Transaction) error {
	conn := GetConn()
	defer conn.Release()

	ctx := context.Background()
	tx, err := conn.Begin(ctx)

	if err != nil {
		return err
	}

	_, err = tx.Exec(
		ctx,
		"INSERT INTO TRANSACTIONS (VALUE, TYPE, DESCRIPTION, CLIENT_ID) VALUES ($1, $2, $3, $4)",
		t.Value,
		t.Type,
		t.Description,
		t.ClientID,
	)

	if err != nil {
		return tx.Rollback(ctx)
	}

	_, err = tx.Exec(
		ctx,
		"UPDATE CLIENTS SET BALANCE = $1 WHERE ID = $2;",
		client.Balance,
		client.ID,
	)

	if err != nil {
		return tx.Rollback(ctx)
	}

	return tx.Commit(ctx)
}
