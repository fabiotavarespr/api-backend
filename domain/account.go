package domain

type Account struct {
	ID          int64   `json:"id"`
	Document    string  `json:"document"`
	CreditLimit float64 `json:"credit_limit"`
}
