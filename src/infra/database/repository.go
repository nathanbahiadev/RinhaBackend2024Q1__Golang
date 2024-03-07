package database

import (
	"context"
	"crebito/src/domain"
	"database/sql"
	"strings"
)

func GetBalanceSQLFunc(ctx context.Context, clientID int32) (*domain.Balance, error) {
	conn := GetConn(ctx)
	defer conn.Release()

	balance := &domain.Balance{
		ClientBalance:    &domain.ClientBalance{},
		LastTransactions: []domain.Transaction{},
	}

	rows, err := conn.Query(ctx, `SELECT * FROM GET_BALANCE($1)`, clientID)

	if err != nil {
		return balance, err
	}

	for rows.Next() {
		var cBalance sql.NullInt32
		var cLimit sql.NullInt32
		var tValue sql.NullInt32
		var tType sql.NullString
		var tDescription sql.NullString
		var tCreatedAt sql.NullString

		if err := rows.Scan(
			&cLimit,
			&cBalance,
			&tValue,
			&tType,
			&tDescription,
			&tCreatedAt,
		); err != nil {
			return nil, err
		}

		balance.ClientBalance.Limit = cLimit.Int32
		balance.ClientBalance.Total = cBalance.Int32

		if tValue.Int32 != 0 {
			balance.LastTransactions = append(balance.LastTransactions, domain.Transaction{
				Value:       tValue.Int32,
				Type:        tType.String,
				Description: tDescription.String,
				CreatedAt:   tCreatedAt.String,
			})
		}
	}

	if balance.ClientBalance.Limit == 0 {
		return balance, domain.ErrClientNotFound
	}

	return balance, nil
}

func CreateTransactionSQLFunc(ctx context.Context, clientID int32, t *domain.Transaction) (*domain.Client, error) {
	conn := GetConn(ctx)
	defer conn.Release()

	row := conn.QueryRow(ctx, `SELECT * FROM CREATE_TRANSACTION($1, $2, $3, $4)`, clientID, t.Value, t.Type, t.Description)
	client := &domain.Client{}

	if err := row.Scan(&client.Balance, &client.AccountLimit); err == nil {
		return client, nil
	} else {
		if strings.Contains(err.Error(), "CLIENT_NOT_FOUND") {
			return client, domain.ErrClientNotFound
		}

		if strings.Contains(err.Error(), "LOW_LIMIT") {
			return client, domain.ErrInsufficientBalance
		}

		return client, err
	}
}
