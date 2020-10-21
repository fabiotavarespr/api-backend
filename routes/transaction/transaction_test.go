package transaction

import (
	"testing"
	"time"

	"github.com/fabiotavarespr/api-backend/domain"
)

func TestIsValidOperationTypes(t *testing.T) {
	var tests = []struct {
		transaction domain.Transaction
		want        bool
	}{
		{domain.Transaction{ID: 1, AccountID: 1, OperationTypeID: 1, Amount: -100.00, EventDT: time.Now()}, true},
		{domain.Transaction{ID: 1, AccountID: 1, OperationTypeID: 1, Amount: 20.00, EventDT: time.Now()}, false},
		{domain.Transaction{ID: 1, AccountID: 1, OperationTypeID: 2, Amount: -30.00, EventDT: time.Now()}, true},
		{domain.Transaction{ID: 1, AccountID: 1, OperationTypeID: 2, Amount: 45.00, EventDT: time.Now()}, false},
		{domain.Transaction{ID: 1, AccountID: 1, OperationTypeID: 3, Amount: -9999.00, EventDT: time.Now()}, true},
		{domain.Transaction{ID: 1, AccountID: 1, OperationTypeID: 3, Amount: 2000.00, EventDT: time.Now()}, false},
		{domain.Transaction{ID: 1, AccountID: 1, OperationTypeID: 4, Amount: 100.00, EventDT: time.Now()}, true},
		{domain.Transaction{ID: 1, AccountID: 1, OperationTypeID: 4, Amount: -90.00, EventDT: time.Now()}, false},
	}
	for _, test := range tests {
		if got := isValidOperationTypes(test.transaction); got != test.want {
			t.Errorf("isValidOperationTypes with Operation %d and amount %g result %t but want %t ", test.transaction.OperationTypeID, test.transaction.Amount, got, test.want)
		}
	}

}
