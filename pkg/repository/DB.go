package repository

import (
	"fmt"
	"testbank/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	usersTable           = "users"
	accountsTable        = "accounts"
	accountsListsTable   = "accounts_list"
	transactionListTable = "transaction_list"
)

func PostgresDBConnect(c *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(c.DBRole, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", c.DBHost, c.DBport, c.DBuser, c.DBname, c.DBpassword, c.DBsslmode))
	logrus.Info(db)
	if err != nil {
		logrus.Fatal("Open: ", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		logrus.Fatal("Ping :", err)
		return nil, err
	}

	logrus.Print("Connection postgres done!")
	return db, nil
}
