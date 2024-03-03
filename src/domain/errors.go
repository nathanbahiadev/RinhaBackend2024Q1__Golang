package domain

import "errors"

var ErrClientNotFound = errors.New("client not found")
var ErrInsufficientBalance = errors.New("insufficient balance")
var ErrInvalidTransactionType = errors.New("invalid transaction type")
var ErrInvalidDescriptionLength = errors.New("description's length must be lower or equal to 10")
var ErrInvalidTransactionValue = errors.New("transaction's values must be greater than 0")

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
	case errors.Is(err, ErrInvalidTransactionType):
		e.Status = 422
	case errors.Is(err, ErrInvalidDescriptionLength):
		e.Status = 422
	case errors.Is(err, ErrInvalidTransactionValue):
		e.Status = 422
	default:
		e.Status = 500
	}

	return e
}
