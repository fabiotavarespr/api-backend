package account

import (
	"github.com/labstack/echo/v4"
)

// Load function to load accounts endpoints
func Load(e *echo.Echo) {
	g := e.Group("/accounts")

	g.POST("", createAccount)
	g.GET("/:account_id", readAccount)
}
