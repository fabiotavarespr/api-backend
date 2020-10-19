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
	logrus.Infof("Starting a new transaction creation process")

	tr := new(domain.Transaction)
	if err := c.Bind(tr); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid transaction data")
	}

	ID, err := database.InsertTransaction(*tr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Error trying to insert a transaction", err),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"ID":      ID,
		"Message": "New transaction successfully created",
	})
}
