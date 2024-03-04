package domain

import (
	"time"
)

type Transaction struct {
	ID          int32     `json:"-"`
	ClientID    int32     `json:"-"`
	Value       int32     `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizada_em"`
}

func (t *Transaction) IsCredit() bool {
	return t.Type == "c"
}

func (t *Transaction) IsDebit() bool {
	return t.Type == "d"
}

func (t *Transaction) Validate() error {
	if !t.IsCredit() && !t.IsDebit() {
		return ErrInvalidTransactionType
	}

	if t.Value <= 0 {
		return ErrInvalidTransactionValue
	}

	if t.Description == "" || len(t.Description) > 10 {
		return ErrInvalidDescriptionLength
	}

	return nil
}

type Client struct {
	ID           int32 `json:"-"`
	AccountLimit int32 `json:"limite"`
	Balance      int32 `json:"saldo"`
}

func (c *Client) AddTransaction(t *Transaction) error {
	if t.IsDebit() {
		nextBalance := c.Balance - t.Value
		if nextBalance < c.AccountLimit*-1 {
			return ErrInsufficientBalance
		}
	}

	if t.IsCredit() {
		c.Balance += t.Value
	} else {
		c.Balance -= t.Value
	}

	return nil
}
