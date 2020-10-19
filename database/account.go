package database

import (
	"github.com/fabiotavarespr/api-backend/domain"
)

// InsertAccount query to insert account data on database
func InsertAccount(documentNumber string, creditLimit float64) (int64, error) {
	res, err := db.Exec("INSERT INTO `account` (`document_number`, `credit_limit`) VALUES (?, ?); ", documentNumber, creditLimit)

	if err == nil {
		id, _ := res.LastInsertId()
		return id, nil
	}

	return 0, err

}

// GetAccount query to return data for a specific account
func GetAccount(id int64) (*domain.Account, error) {

	account := new(domain.Account)

	err := db.QueryRow("SELECT `id`, `document_number`, `credit_limit` FROM `account` WHERE `id`=?", id).
		Scan(&account.ID, &account.Document, &account.CreditLimit)

	return account, err
}

// UpdateAccountCreditLimit query to update account
func UpdateAccountCreditLimit(idAccount int64, newLimit float64) (int64, error) {
	res, err := db.Exec("UPDATE `account` SET `credit_limit` = ? WHERE `id` = ?;", newLimit, idAccount)

	if err != nil {
		return 0, err
	} else {
		return res.RowsAffected()
	}
}
