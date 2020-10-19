package transaction

import (
	"github.com/labstack/echo/v4"
)

// Load function to load transactions endpoints
func Load(e *echo.Echo) {
	transactions := e.Group("/transactions")

	transactions.POST("", createTransaction)
}
