package routes

import (
	"github.com/fabiotavarespr/api-backend/routes/account"
	"github.com/fabiotavarespr/api-backend/routes/transaction"
	"github.com/labstack/echo/v4"
)

//Register routes
func Register(e *echo.Echo) {
	account.Load(e)
	transaction.Load(e)
}
