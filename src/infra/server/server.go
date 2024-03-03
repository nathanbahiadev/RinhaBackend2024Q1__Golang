package server

import (
	"crebito/src/domain"
	"crebito/src/infra/server/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Port       string
	Router     *chi.Mux
	Repository domain.RepositoryInterface
}

func New(port string, repository domain.RepositoryInterface) Server {
	server := Server{
		Port:       port,
		Repository: repository,
		Router:     chi.NewRouter(),
	}

	server.Router.Use(middleware.Logger)

	server.Router.Post(
		"/clientes/{id}/transacoes",
		handlers.NewCreateTransactionHandler(
			domain.NewCreateTransactionUseCase(
				server.Repository,
			),
		).Handle,
	)

	server.Router.Get(
		"/clientes/{id}/extrato",
		handlers.NewGetBalanceHandler(
			domain.NewGetBalanceUseCase(
				server.Repository,
			),
		).Handle,
	)

	return server
}

func (server Server) Start() error {
	return http.ListenAndServe(server.Port, server.Router)
}
