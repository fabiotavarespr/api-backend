package transaction

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/fabiotavarespr/api-backend/database"
	"github.com/fabiotavarespr/api-backend/domain"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func createTransaction(c echo.Context) error {
	logrus.Infof("Starting createTransaction process")

	tr := new(domain.Transaction)
	if err := c.Bind(tr); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid transaction data")
	}

	if !isValidOperationTypes(*tr) {
		return c.JSON(http.StatusBadRequest, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid value with payment"),
		})
	}

	account, err := database.GetAccount(tr.AccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			logrus.Errorf("account_id %v doesn't exist", tr.AccountID)
			return c.JSON(http.StatusNotFound, &echo.HTTPError{
				Code:    http.StatusNotFound,
				Message: fmt.Sprintf("account_id %v doesn't exist", tr.AccountID),
			})
		}
	}

	if !isValidCreditLimit(account.CreditLimit, tr.Amount) {
		return c.JSON(http.StatusBadRequest, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid credit limit to payment"),
		})
	}

	ID, err := database.InsertTransaction(*tr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Error trying to insert a transaction"),
		})
	}

	_, _ = database.UpdateAccountCreditLimit(tr.AccountID, calculateCreditLimit(account.CreditLimit, tr.Amount))

	logrus.Infof("Ending createTransaction process")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"ID":      ID,
		"Message": "New transaction successfully created",
	})
}

func isValidOperationTypes(transaction domain.Transaction) bool {

	if transaction.OperationTypeID == domain.CompraAVista ||
		transaction.OperationTypeID == domain.CompraParcelada ||
		transaction.OperationTypeID == domain.Saque {
		if transaction.Amount > 0 {
			return false
		}
	}

	if transaction.OperationTypeID == domain.Pagamento {
		if transaction.Amount < 0 {
			return false
		}
	}

	return true
}

func isValidCreditLimit(creditLimit, amount float64) bool {
	limit := calculateCreditLimit(creditLimit, amount)
	if limit < 0 {
		return false
	}

	return true
}

func calculateCreditLimit(creditLimit, amount float64) float64 {
	return creditLimit + amount
}
