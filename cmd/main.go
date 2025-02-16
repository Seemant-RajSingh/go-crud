package main

import (
	"database/sql"
	"log"

	"github.com/Seemant-RajSingh/go-crud/cmd/api"
	"github.com/Seemant-RajSingh/go-crud/config"
	db "github.com/Seemant-RajSingh/go-crud/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser, // fields User, Passwd taken from dsn.go(lib) Config struct
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal()
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db) // db is a pointer to sql.DB
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) { // pointer to sql database
	err := db.Ping() // CONNECTING TO DB
	if err != nil {
		log.Fatal()
	}

	log.Println("DB: Succesfully connected")
}
