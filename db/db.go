package db

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var once sync.Once

func connect() {
    connectionString := os.Getenv("PG_CONNECTION_STRING")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

    err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db
}

func GetDB() *sql.DB {
	once.Do(connect)

	return DB
}

