package usecases

import (
	"context"
	"crebito/src/domain"
)

type InputCreateTransactionUseCase struct {
	ClientID    int32
	Value       int32  `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

type TCreateTransactionRepoFunc func(ctx context.Context, clientID int32, t *domain.Transaction) (*domain.Client, error)
type TCreateTransactionUseCaseFunc func(ctx context.Context, input InputCreateTransactionUseCase, createFunc TCreateTransactionRepoFunc) (*domain.Client, *domain.Exception)

func CreateTransactionUseCase(
	ctx context.Context,
	input InputCreateTransactionUseCase,
	createFunc TCreateTransactionRepoFunc,
) (*domain.Client, *domain.Exception) {

	transaction := &domain.Transaction{
		Value:       input.Value,
		Type:        input.Type,
		Description: input.Description,
	}

	if !transaction.IsValid() {
		return nil, domain.HandleError(domain.ErrInvalidTransaction)
	}

	client, err := createFunc(ctx, input.ClientID, transaction)

	if err != nil {
		return client, domain.HandleError(err)
	}

	return client, nil
}
