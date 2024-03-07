package server

import (
	"context"
	"crebito/src/domain/usecases"
	"crebito/src/infra/server/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Port                         string
	Context                      context.Context
	Router                       *chi.Mux
	CreateTransactionRepoFunc    usecases.TCreateTransactionRepoFunc
	CreateTransactionUseCaseFunc usecases.TCreateTransactionUseCaseFunc
	GetBalanceRepoFunc           usecases.TGetBalanceRepoFunc
	GetBalanceUseCaseFunc        usecases.TGetBalanceUseCaseFunc
}

func (server Server) Start() error {
	server.Router = chi.NewRouter()

	server.Router.Post(
		"/clientes/{id}/transacoes",
		handlers.CreateTransactionHandler{
			Context:    server.Context,
			UseCase:    server.CreateTransactionUseCaseFunc,
			CreateFunc: server.CreateTransactionRepoFunc,
		}.Handle,
	)

	server.Router.Get(
		"/clientes/{id}/extrato",
		handlers.GetBalanceHandler{
			Context:        server.Context,
			UseCase:        server.GetBalanceUseCaseFunc,
			GetBalanceFunc: server.GetBalanceRepoFunc,
		}.Handle,
	)

	return http.ListenAndServe(server.Port, server.Router)
}
