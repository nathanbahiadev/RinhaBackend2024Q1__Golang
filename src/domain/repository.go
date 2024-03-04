package domain

type RepositoryInterface interface {
	GetClient(clientID int32) (*Client, error)
	CreateTransactionAndUpdateBalance(client *Client, t *Transaction) error
	FindLastTransactionsByClient(clientID int32) ([]Transaction, error)
}

type CacheRepositoryInterface interface {
	GetClient(clientID int32) (*Client, error)
	UpdateClient(client *Client) error
}
