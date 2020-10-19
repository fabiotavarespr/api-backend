package database

import (
	"fmt"

	"github.com/fabiotavarespr/api-backend/infra/env"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var (
	db  *sqlx.DB
	err error
)

// Config - opens a connection to database
func Config() {

	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8",
		env.Get().DatabaseUsername,
		env.Get().DatabasePassword,
		env.Get().DatabaseHostname,
		env.Get().DatabasePort,
		env.Get().DatabaseName)

	db, err = sqlx.Open("mysql", dbConnection)
	if err != nil {
		logrus.Panicln("Error connecting to the database", err)
		return
	}
	logrus.Info("Database connection is working properly!")
}
