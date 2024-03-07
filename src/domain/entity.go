package domain

type Transaction struct {
	Value       int32  `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
	CreatedAt   string `json:"realizada_em"`
}

func (t *Transaction) IsValid() bool {
	if t.Type != "c" && t.Type != "d" {
		return false
	}

	if t.Value <= 0 {
		return false
	}

	if len(t.Description) == 0 || len(t.Description) > 10 {
		return false
	}

	return true
}

type Client struct {
	AccountLimit int32 `json:"limite"`
	Balance      int32 `json:"saldo"`
}

type ClientBalance struct {
	Limit int32 `json:"limite"`
	Total int32 `json:"total"`
}

type Balance struct {
	ClientBalance    *ClientBalance `json:"saldo"`
	LastTransactions []Transaction  `json:"ultimas_transacoes"`
}
