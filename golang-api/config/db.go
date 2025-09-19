package config

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var DB *sql.DB

func ConnectDB() {
    var err error
    connStr := "host=localhost port=5433 user=postgres password='' dbname=ticketing2_db sslmode=disable"
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to PostgreSQL")
}
