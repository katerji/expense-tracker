package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/katerji/expense-tracker/db/generated"
	"github.com/katerji/expense-tracker/env"
	"time"
)

type client struct {
	*generated.Queries
}

var c *client

func getInstance() *client {
	if c != nil {
		return c
	}
	conn, err := initConnection()
	if err != nil {
		panic(err)
	}
	c = &client{
		generated.New(conn),
	}

	return c
}

func initConnection() (*sql.DB, error) {
	dbHost := env.DbHost()
	dbUser := env.DbUser()
	dbPort := env.DbPort()
	dbPass := env.DbPassword()
	dbName := env.DbName()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
