package database

import (
	"github.com/fabiotavarespr/api-backend/domain"
)

// InsertAccount query to insert account data on database
func InsertAccount(documentNumber string) (int64, error) {
	res, err := db.Exec("INSERT INTO `account` (`document_number`) VALUES (?); ", documentNumber)

	if err == nil {
		id, _ := res.LastInsertId()
		return id, nil
	}

	return 0, err

}

// GetAccount query to return data for a specific account
func GetAccount(id int64) (*domain.Account, error) {

	account := new(domain.Account)

	err := db.QueryRow("SELECT `id`, `document_number` FROM `account` WHERE `id`=?", id).
		Scan(&account.ID, &account.Document)

	return account, err
}
