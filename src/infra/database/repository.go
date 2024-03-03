package database

import (
	"context"
	"crebito/src/domain"
	"crebito/src/infra/database/sqlc"
	"database/sql"
)

type SqlRepository struct {
	DB      *sql.DB
	Queries *sqlc.Queries
}

func NewSqlRepositories(db *sql.DB) SqlRepository {
	return SqlRepository{DB: db, Queries: sqlc.New(db)}
}

func (r *SqlRepository) GetClient(clientID int32) (*domain.Client, error) {
	if client, err := r.Queries.GetClient(context.Background(), clientID); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrClientNotFound
		}
		return nil, err
	} else {
		return &domain.Client{
			ID:           client.ID,
			AccountLimit: client.AccountLimit,
			Balance:      client.Balance,
		}, nil
	}
}

func (r *SqlRepository) FindLastTransactionsByClient(clientID int32) ([]domain.Transaction, error) {
	transactions, err := r.Queries.FindLastTransactionsByClient(context.Background(), clientID)

	if err != nil {
		return nil, err
	}

	var result []domain.Transaction

	for _, t := range transactions {
		result = append(result, domain.Transaction{
			Value:       t.Value,
			Type:        t.Type,
			CreatedAt:   t.CreatedAt.Time,
			Description: t.Description,
		})
	}

	return result, nil
}

func (r *SqlRepository) CreateTransactionAndUpdateBalance(client *domain.Client, t *domain.Transaction) error {
	tx, err := r.DB.Begin()

	if err != nil {
		return err
	}

	ctx := context.Background()
	qtx := r.Queries.WithTx(tx)

	if err := qtx.CreateTransaction(ctx, sqlc.CreateTransactionParams{
		Value:       t.Value,
		Type:        t.Type,
		ClientID:    t.ClientID,
		Description: t.Description,
	}); err != nil {
		return tx.Rollback()
	}

	if err := qtx.UpdateBalance(ctx, sqlc.UpdateBalanceParams{
		ID:      client.ID,
		Balance: client.Balance,
	}); err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}
