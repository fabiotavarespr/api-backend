package account

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/fabiotavarespr/api-backend/database"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func createAccount(c echo.Context) error {
	logrus.Infof("Starting createAccount process")

	account := new(struct {
		DocNumber   string  `json:"document_number"`
		CreditLimit float64 `json:"credit_limit"`
	})

	if err := c.Bind(account); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid account data")
	}

	ID, err := database.InsertAccount(account.DocNumber, account.CreditLimit)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return c.JSON(http.StatusBadRequest, &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("Account with document %v already exists", account.DocNumber),
			})
		}
		if strings.Contains(err.Error(), "Data too long for column 'document_number'") {
			return c.JSON(http.StatusBadRequest, &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("%v is an invalid document, document with a maximum of 14 characters ", account.DocNumber),
			})
		}
		return c.JSON(http.StatusBadRequest, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Error trying to insert a account data", err),
		})
	}

	logrus.Infof("Ending createAccount process")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"ID":      ID,
		"Message": "New account successfully created",
	})
}

func readAccount(c echo.Context) error {
	logrus.Infof("Starting readAccount with id %v", c.Param("account_id"))
	accountID, err := strconv.ParseInt(c.Param("account_id"), 10, 64)
	if err != nil {
		logrus.Errorf("account_id must be a number: %v", c.Param("account_id"))
		return c.JSON(http.StatusBadRequest, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("account_id must be a number: %v", c.Param("account_id")),
		})
	}

	accountResponse, err := database.GetAccount(accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			logrus.Errorf("account_id %v doesn't exist", accountID)
			return c.JSON(http.StatusNotFound, &echo.HTTPError{
				Code:    http.StatusNotFound,
				Message: fmt.Sprintf("account_id %v doesn't exist", accountID),
			})
		}
	}
	logrus.Infof("Ending readAccount with id %v", accountID)
	return c.JSON(http.StatusOK, accountResponse)
}
