package main

import (
	"crebito/src/infra/database"
	"crebito/src/infra/server"
	"log"
	"os"
)

func main() {
	err := database.New(os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("Falha ao iniciar banco de dados: %s", err)
	}

	httpServer := server.New(os.Getenv("PORT"), database.NewSqlRepositories())

	if err := httpServer.Start(); err != nil {
		log.Fatalf("Falha ao iniciar servidor http: %s", err)
	}
}
