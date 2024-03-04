package usecases

import (
	"crebito/src/domain"
	"time"
)

type GetBalanceUseCase struct {
	Repository domain.RepositoryInterface
}

func NewGetBalanceUseCase(repository domain.RepositoryInterface) GetBalanceUseCase {
	return GetBalanceUseCase{
		Repository: repository,
	}
}

type OutputGetBalanceUseCase struct {
	Balance struct {
		Total int32     `json:"total"`
		Date  time.Time `json:"data_extrato"`
		Limit int32     `json:"limite"`
	} `json:"saldo"`
	LastTransactions []domain.Transaction `json:"ultimas_transacoes"`
}

func (useCase GetBalanceUseCase) Execute(clientID int32) (OutputGetBalanceUseCase, *domain.Exception) {
	var output OutputGetBalanceUseCase

	client, err := useCase.Repository.GetClient(clientID)

	if err != nil {
		return output, domain.HandleError(err)
	}

	transactions, err := useCase.Repository.FindLastTransactionsByClient(clientID)

	if err != nil {
		return output, domain.HandleError(err)
	}

	output.Balance.Date = time.Now()
	output.Balance.Limit = client.AccountLimit
	output.Balance.Total = client.Balance
	output.LastTransactions = transactions

	return output, nil
}
