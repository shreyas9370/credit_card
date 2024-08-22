package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Shreyas9370@@tcp(localhost:3306)/creditcard")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
