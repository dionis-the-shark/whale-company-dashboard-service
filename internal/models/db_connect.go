package models

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DBConnect() {
	user := "root"
	pass := "testpassword"

	cfg := mysql.Config{
		User:   user,
		Passwd: pass,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "whale_products",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pringErr := db.Ping()
	if pringErr != nil {
		log.Fatal(pringErr)
	}
}
