package account

import (
	"github.com/labstack/echo/v4"
)

// Load function to load accounts endpoints
func Load(e *echo.Echo) {
	account := e.Group("/accounts")

	account.POST("", createAccount)
	account.GET("/:account_id", readAccount)
}
