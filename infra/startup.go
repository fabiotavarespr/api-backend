package infra

import (
	"github.com/fabiotavarespr/api-backend/database"
	"github.com/fabiotavarespr/api-backend/infra/env"
	"github.com/fabiotavarespr/api-backend/infra/server"
	"github.com/sirupsen/logrus"
)

// Config is a function
func Config() {
	configLogFormatter()

	logrus.Info("Starting environment and log configuration...")
	env.Config()

	logrus.Info("Starting database connection...")
	database.Config()

	logrus.Info("Starting HTTP server...")
	server.Config()

}

func configLogFormatter() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.999999",
	})
}
