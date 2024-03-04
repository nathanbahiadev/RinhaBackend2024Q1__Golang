package domain_test

import (
	"crebito/src/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTransactions(t *testing.T) {
	client := domain.Client{ID: 1, AccountLimit: 100000, Balance: 0}
	transaction := domain.Transaction{ID: 1, Value: 1, Description: "transaction", Type: "c"}

	for i := 0; i < 100; i++ {
		err := client.AddTransaction(&transaction)
		assert.Nil(t, err)
	}

	assert.Equal(t, client.Balance, int32(100))

	transaction.Type = "d"

	for i := 0; i < 100; i++ {
		err := client.AddTransaction(&transaction)
		assert.Nil(t, err)
	}

	assert.Equal(t, client.Balance, int32(0))
}

func TestAddTransactionsWithInsufficientBalance(t *testing.T) {
	client := domain.Client{ID: 1, AccountLimit: 1000, Balance: 0}
	transaction := domain.Transaction{ID: 1, Value: 500, Description: "transaction", Type: "d"}

	err := client.AddTransaction(&transaction)
	assert.Nil(t, err)
	assert.Equal(t, client.Balance, int32(-500))

	err = client.AddTransaction(&transaction)
	assert.Nil(t, err)
	assert.Equal(t, client.Balance, int32(-1000))

	err = client.AddTransaction(&transaction)
	assert.Error(t, err)
	assert.Equal(t, client.Balance, int32(-1000))
}
