package domain

type RepositoryInterface interface {
	GetClient(clientID int32) (*Client, error)
	CreateTransactionAndUpdateBalance(clientID int32, t *Transaction) (*Client, error)
	FindLastTransactionsByClient(clientID int32) ([]Transaction, error)
}
