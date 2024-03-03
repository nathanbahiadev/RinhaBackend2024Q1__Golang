package handlers

import (
	"crebito/src/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CreateTransactionHandler struct {
	UseCase domain.CreateTransactionUseCase
}

func NewCreateTransactionHandler(useCase domain.CreateTransactionUseCase) CreateTransactionHandler {
	return CreateTransactionHandler{UseCase: useCase}
}

func (h CreateTransactionHandler) Handle(w http.ResponseWriter, r *http.Request) {
	clientID, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	var input domain.InputCreateTransactionUseCase
	input.ClientID = int32(clientID)

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	output, exception := h.UseCase.Execute(input)

	if exception != nil {
		w.WriteHeader(exception.Status)
		json.NewEncoder(w).Encode(exception)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(output)
}
