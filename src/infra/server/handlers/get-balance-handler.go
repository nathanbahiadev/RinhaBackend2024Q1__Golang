package handlers

import (
	"context"
	"crebito/src/domain/usecases"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type GetBalanceHandler struct {
	Context        context.Context
	UseCase        usecases.TGetBalanceUseCaseFunc
	GetBalanceFunc usecases.TGetBalanceRepoFunc
}

func (handler GetBalanceHandler) Handle(w http.ResponseWriter, r *http.Request) {

	clientID, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		w.WriteHeader(400)
		return
	}

	output, exception := handler.UseCase(handler.Context, int32(clientID), handler.GetBalanceFunc)

	if exception != nil {
		w.WriteHeader(exception.Status)
		json.NewEncoder(w).Encode(exception)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(output)
}
