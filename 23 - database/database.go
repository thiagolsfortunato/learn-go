package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	urlConn := "learn-go:learn-go@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, error := sql.Open("mysql", urlConn)
	if error != nil {
		log.Fatal(error)
	}
	defer db.Close()

	if error = db.Ping(); error != nil {
		log.Fatal(error)
	}

	fmt.Println("Connection open!")

	result, error := db.Query("select * from users")
	if error != nil {
		log.Fatal(error)
	}
	defer result.Close()

	fmt.Println(result)
}
