package main

import (
	"context"
	"crebito/src/domain"
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

	repository := database.NewSqlRepositories(ctx)
	httpServer := server.New(os.Getenv("PORT"), repository)

	startEngines(repository)
	err = httpServer.Start()

	if err != nil {
		log.Fatalf("Falha ao iniciar servidor http: %s", err)
	}
}

func startEngines(repository domain.RepositoryInterface) {
	for i := 0; i < 50; i++ {
		repository.GetClient(int32(i))
	}
}
