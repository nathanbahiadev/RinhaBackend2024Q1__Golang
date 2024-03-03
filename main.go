package main

import (
	"crebito/src/infra/database"
	"crebito/src/infra/server"
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("Iniciando aplicação...")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	log.Println(dsn)

	db, err := database.New(dsn)
	if err != nil {
		log.Fatalf("Falha ao iniciar banco de dados: %s", err)
	}

	log.Println("Banco de dados iniciado com sucesso!")

	repository := database.NewSqlRepositories(db)
	httpServer := server.New(fmt.Sprintf(":%s", os.Getenv("PORT")), &repository)

	log.Println("Iniciando servidor http...")

	if err := httpServer.Start(); err != nil {
		log.Fatalf("Falha ao iniciar servidor http: %s", err)
	}
}
