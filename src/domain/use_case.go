package domain

import (
	"time"
)

type CreateTransactionUseCase struct {
	Repository RepositoryInterface
}

func NewCreateTransactionUseCase(repository RepositoryInterface) CreateTransactionUseCase {
	return CreateTransactionUseCase{
		Repository: repository,
	}
}

type InputCreateTransactionUseCase struct {
	ClientID    int32
	Value       int32  `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

type OutputCreateTransactionUseCase struct {
	Limit   int32 `json:"limite"`
	Balance int32 `json:"saldo"`
}

func (useCase CreateTransactionUseCase) Execute(input InputCreateTransactionUseCase) (OutputCreateTransactionUseCase, *Exception) {
	var output OutputCreateTransactionUseCase

	transaction := &Transaction{
		ClientID:    input.ClientID,
		Value:       input.Value,
		Type:        input.Type,
		Description: input.Description,
	}

	if err := transaction.Validate(); err != nil {
		return output, HandleError(err)
	}

	client, err := useCase.Repository.GetClient(input.ClientID)

	if err != nil {
		return output, HandleError(err)
	}

	if err := client.AddTransaction(transaction); err != nil {
		return output, HandleError(err)
	}

	if err := useCase.Repository.CreateTransactionAndUpdateBalance(client, transaction); err != nil {
		return output, HandleError(err)
	}

	output.Balance = client.Balance
	output.Limit = client.AccountLimit

	return output, nil
}

type GetBalanceUseCase struct {
	Repository      RepositoryInterface
	CacheRepository CacheRepositoryInterface
}

func NewGetBalanceUseCase(repository RepositoryInterface) GetBalanceUseCase {
	return GetBalanceUseCase{
		Repository: repository,
	}
}

type OutputGetBalanceUseCase struct {
	LastTransactions []Transaction `json:"ultimas_transacoes"`
	Balance          struct {
		Total int32     `json:"total"`
		Date  time.Time `json:"data_extrato"`
		Limit int32     `json:"limite"`
	} `json:"saldo"`
}

func (useCase GetBalanceUseCase) Execute(clientID int32) (OutputGetBalanceUseCase, *Exception) {
	var output OutputGetBalanceUseCase

	client, err := useCase.Repository.GetClient(clientID)

	if err != nil {
		return output, HandleError(err)
	}

	transactions, err := useCase.Repository.FindLastTransactionsByClient(clientID)

	if err != nil {
		return output, HandleError(err)
	}

	output.Balance.Date = time.Now()
	output.Balance.Limit = client.AccountLimit
	output.Balance.Total = client.Balance
	output.LastTransactions = transactions

	return output, nil
}
