package domain

import "errors"

var ErrClientNotFound = errors.New("client not found")
var ErrInsufficientBalance = errors.New("insufficient balance")
var ErrInvalidTransaction = errors.New("invalid transaction")

type Exception struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func HandleError(err error) *Exception {
	e := &Exception{Message: err.Error()}

	switch true {
	case errors.Is(err, ErrClientNotFound):
		e.Status = 404
	case errors.Is(err, ErrInsufficientBalance):
		e.Status = 422
	case errors.Is(err, ErrInvalidTransaction):
		e.Status = 422
	default:
		e.Status = 500
	}

	return e
}
