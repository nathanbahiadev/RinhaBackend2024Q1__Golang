package handlers

import (
	"context"
	"crebito/src/domain/usecases"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CreateTransactionHandler struct {
	Context    context.Context
	UseCase    usecases.TCreateTransactionUseCaseFunc
	CreateFunc usecases.TCreateTransactionRepoFunc
}

func (handler CreateTransactionHandler) Handle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	clientID, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		w.WriteHeader(400)
		return
	}

	var input usecases.InputCreateTransactionUseCase
	input.ClientID = int32(clientID)

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(400)
		return
	}

	output, exception := handler.UseCase(handler.Context, input, handler.CreateFunc)

	if exception != nil {
		w.WriteHeader(exception.Status)
		json.NewEncoder(w).Encode(exception)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(output)
}
