package server

import (
	"github.com/fabiotavarespr/api-backend/infra/env"
	"github.com/fabiotavarespr/api-backend/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

// Config is a function
func Config() {

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	routes.Register(e)

	logrus.Fatal(e.Start(env.Get().ServerHostname + ":" + env.Get().ServerPort))

}
