package main

import (
	"context"
	"crebito/src/domain/usecases"
	"crebito/src/infra/database"
	"crebito/src/infra/server"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	db, err := database.New(ctx, os.Getenv("DATABASE_URL"))
	defer db.Close()

	if err != nil {
		log.Fatalf("Falha ao iniciar banco de dados: %s", err)
	}

	httpServer := server.Server{
		Port:                         os.Getenv("PORT"),
		Context:                      ctx,
		CreateTransactionUseCaseFunc: usecases.CreateTransactionUseCase,
		GetBalanceUseCaseFunc:        usecases.GetBalanceUseCase,
		CreateTransactionRepoFunc:    database.CreateTransactionSQLFunc,
		GetBalanceRepoFunc:           database.GetBalanceSQLFunc,
	}

	err = httpServer.Start()

	if err != nil {
		log.Fatalf("Falha ao iniciar servidor http: %s", err)
	}
}
