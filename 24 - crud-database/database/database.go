package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // drive mysql connection
)

func Connect() (*sql.DB, error) {
	urlConn := "learn-go:learn-go@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, error := sql.Open("mysql", urlConn)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		return nil, error
	}

	return db, nil
}
