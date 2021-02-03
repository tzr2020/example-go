package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

const (
	server   = "localhost"
	port     = "3306"
	user     = "root"
	password = "123456"
	database = "test"
)

func main() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, password, server, port, database)

	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := context.Background()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("database connected!")
}
