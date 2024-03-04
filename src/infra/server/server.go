package server

import (
	"crebito/src/domain"
	"crebito/src/infra/server/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Port       string
	Router     *chi.Mux
	Repository domain.RepositoryInterface
	// CacheRepository domain.CacheRepositoryInterface
}

func New(
	port string,
	repository domain.RepositoryInterface,
	// cacheRepository domain.CacheRepositoryInterface,
) Server {
	server := Server{
		Port:       port,
		Repository: repository,
		// CacheRepository: cacheRepository,
		Router: chi.NewRouter(),
	}

	server.Router.Post(
		"/clientes/{id}/transacoes",
		handlers.NewCreateTransactionHandler(
			domain.NewCreateTransactionUseCase(
				server.Repository,
				// server.CacheRepository,
			),
		).Handle,
	)

	server.Router.Get(
		"/clientes/{id}/extrato",
		handlers.NewGetBalanceHandler(
			domain.NewGetBalanceUseCase(
				server.Repository,
				// server.CacheRepository,
			),
		).Handle,
	)

	return server
}

func (server Server) Start() error {
	return http.ListenAndServe(server.Port, server.Router)
}
