package server

import (
	"crebito/src/domain"
	"crebito/src/domain/usecases"
	"crebito/src/infra/server/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Port       string
	Router     *chi.Mux
	Repository domain.RepositoryInterface
}

func (server Server) Start() error {
	return http.ListenAndServe(server.Port, server.Router)
}

func New(
	port string,
	repository domain.RepositoryInterface,
) Server {
	server := Server{
		Port:       port,
		Repository: repository,
		Router:     chi.NewRouter(),
	}

	server.Router.Post(
		"/clientes/{id}/transacoes",
		handlers.NewCreateTransactionHandler(usecases.NewCreateTransactionUseCase(server.Repository)).Handle,
	)

	server.Router.Get(
		"/clientes/{id}/extrato",
		handlers.NewGetBalanceHandler(usecases.NewGetBalanceUseCase(server.Repository)).Handle,
	)

	return server
}
