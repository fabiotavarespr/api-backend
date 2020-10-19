package env

import (
	"flag"

	"github.com/fabiotavarespr/api-backend/infra/log"
	"github.com/sirupsen/logrus"
)

// Environment Environment
type Environment struct {
	ServerHostname   string
	ServerPort       string
	DatabaseHostname string
	DatabasePort     int
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
}

type config struct {
	DateFormat string
}

var environment Environment

// Config - Configuration for environment variables
func Config() {

	logLevel := flag.String("log-level", "info", "debug, info, warning, error")
	serverHostname := flag.String("server-hostname", "0.0.0.0", "The server hostname")
	serverPort := flag.String("server-port", "8080", "The server port")
	databaseHostname := flag.String("database-hostname", "127.0.0.1", "The database hostname.")
	databasePort := flag.Int("database-port", 3306, "The database port.")
	databaseUsername := flag.String("database-username", "root", "The database username.")
	databasePassword := flag.String("database-password", "passwd123", "The database password.")
	databaseName := flag.String("database-name", "backend", "The database name.")

	flag.Parse()

	log.ConfigLog(logLevel)
	logrus.WithFields(logrus.Fields{"log-level": *logLevel}).Debug("Log level defined")

	environment.ServerHostname = *serverHostname
	logrus.WithFields(logrus.Fields{"server-hostname": *serverHostname}).Debug("The server hostname defined")

	environment.ServerPort = *serverPort
	logrus.WithFields(logrus.Fields{"server-port": *serverPort}).Debug("The server port defined")

	environment.DatabaseHostname = *databaseHostname
	logrus.WithFields(logrus.Fields{"database-hostname": *databaseHostname}).Debug("The database hostname defined")

	environment.DatabasePort = *databasePort
	logrus.WithFields(logrus.Fields{"database-port": *databasePort}).Debug("The database port defined")

	environment.DatabaseUsername = *databaseUsername
	logrus.WithFields(logrus.Fields{"database-username": *databaseUsername}).Debug("The database username defined")

	environment.DatabasePassword = *databasePassword
	logrus.WithFields(logrus.Fields{"database-password": *databasePassword}).Debug("The database password defined")

	environment.DatabaseName = *databaseName
	logrus.WithFields(logrus.Fields{"database-name": *databaseName}).Debug("The database name defined")

}

// Get env from external
func Get() Environment {
	return environment
}
