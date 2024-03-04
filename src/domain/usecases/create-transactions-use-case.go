package usecases

import (
	"crebito/src/domain"
)

type CreateTransactionUseCase struct {
	Repository domain.RepositoryInterface
}

func NewCreateTransactionUseCase(repository domain.RepositoryInterface) CreateTransactionUseCase {
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

func (useCase CreateTransactionUseCase) Execute(input InputCreateTransactionUseCase) (OutputCreateTransactionUseCase, *domain.Exception) {
	var output OutputCreateTransactionUseCase

	transaction := &domain.Transaction{
		ClientID:    input.ClientID,
		Value:       input.Value,
		Type:        input.Type,
		Description: input.Description,
	}

	if err := transaction.Validate(); err != nil {
		return output, domain.HandleError(err)
	}

	client, err := useCase.Repository.CreateTransactionAndUpdateBalance(input.ClientID, transaction)

	if err != nil {
		return output, domain.HandleError(err)
	}

	output.Balance = client.Balance
	output.Limit = client.AccountLimit

	return output, nil
}
