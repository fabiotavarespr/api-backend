package transaction

import (
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

	if !validationOperationTypes(*tr) {
		return c.JSON(http.StatusBadRequest, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid value with payment"),
		})
	}

	validCredit, newLimit := validationCreditLimit(*tr)

	if !validCredit {
		return c.JSON(http.StatusBadRequest, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid credit limit to payment"),
		})
	}

	ID, err := database.InsertTransaction(*tr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Error trying to insert a transaction", err),
		})
	}

	_, _ = database.UpdateAccountCreditLimit(tr.AccountID, newLimit)

	logrus.Infof("Ending createTransaction process")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"ID":      ID,
		"Message": "New transaction successfully created",
	})
}

func validationOperationTypes(transaction domain.Transaction) bool {

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

func validationCreditLimit(transaction domain.Transaction) (bool, float64) {
	account, err := database.GetAccount(transaction.AccountID)
	if err != nil {
		return false, 0
	}

	limit := account.CreditLimit + transaction.Amount

	if limit < 0 {
		return false, 0
	}

	return true, limit
}
