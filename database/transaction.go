package database

import (
	"time"

	"github.com/fabiotavarespr/api-backend/domain"
)

// InsertTransaction query to insert transaction data on database
func InsertTransaction(transaction domain.Transaction) (int64, error) {
	res, err := db.Exec("INSERT INTO `transaction` (`account_id`,`operation_type_id`,`amount`, `event_date`) VALUES (?,?,?,?);", transaction.AccountID, transaction.OperationTypeID, transaction.Amount, time.Now().Format("2006-01-02 15:04:05.999999"))

	if err == nil {
		id, _ := res.LastInsertId()
		return id, nil
	}

	return 0, err
}
