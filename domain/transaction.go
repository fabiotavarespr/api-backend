package domain

import "time"

type OperationType int

const (
	CompraAVista OperationType = 1 + iota
	CompraParcelada
	Saque
	Pagamento
)

var operationTypes = []string{
	"COMPRA A VISTA",
	"COMPRA PARCELADA",
	"SAQUE",
	"PAGAMENTO",
}

func (t OperationType) String() string {
	return operationTypes[t-1]
}

type Transaction struct {
	ID              int64         `json:"id"`
	AccountID       int64         `json:"account_id"`
	OperationTypeID OperationType `json:"operation_type_id"`
	Amount          float64       `json:"amount"`
	EventDT         time.Time     `json:"event_dt"`
}
