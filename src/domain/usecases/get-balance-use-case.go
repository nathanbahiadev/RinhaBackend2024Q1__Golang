package usecases

import (
	"context"
	"crebito/src/domain"
)

type TGetBalanceRepoFunc func(ctx context.Context, clientID int32) (*domain.Balance, error)
type TGetBalanceUseCaseFunc func(ctx context.Context, clientID int32, getBalanceFunc TGetBalanceRepoFunc) (*domain.Balance, *domain.Exception)

func GetBalanceUseCase(ctx context.Context, clientID int32, getBalanceFunc TGetBalanceRepoFunc) (*domain.Balance, *domain.Exception) {

	output, err := getBalanceFunc(ctx, clientID)

	if err != nil {
		return output, domain.HandleError(err)
	}

	return output, nil
}
