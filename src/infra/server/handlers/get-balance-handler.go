package handlers

import (
	"crebito/src/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type GetBalanceHandler struct {
	UseCase domain.GetBalanceUseCase
}

func NewGetBalanceHandler(useCase domain.GetBalanceUseCase) GetBalanceHandler {
	return GetBalanceHandler{UseCase: useCase}
}

func (h GetBalanceHandler) Handle(w http.ResponseWriter, r *http.Request) {
	clientID, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	output, exception := h.UseCase.Execute(int32(clientID))

	if exception != nil {
		w.WriteHeader(exception.Status)
		json.NewEncoder(w).Encode(exception)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(output)
}
