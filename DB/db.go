package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

// essentially returns db struct (pointer to it)
func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN()) // db is a type of struct returned
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
