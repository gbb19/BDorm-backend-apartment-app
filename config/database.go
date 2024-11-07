
package config

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
    var err error
    dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    if err := DB.Ping(); err != nil {
        log.Fatal(err)
    }

    log.Println("Database connected")
}
