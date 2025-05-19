package internal

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(dataSourceName string) {
    var err error
    DB, err = sql.Open("sqlite3", dataSourceName)
    if err != nil {
        log.Fatal("Cannot open database:", err)
    }

    createSubscriptionsTable := `
    CREATE TABLE IF NOT EXISTS subscriptions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT NOT NULL UNIQUE,
        city TEXT NOT NULL,
        frequency TEXT NOT NULL,
        confirmed BOOLEAN DEFAULT 0,
        token TEXT
    );
    `
    _, err = DB.Exec(createSubscriptionsTable)
    if err != nil {
        log.Fatal("Cannot create subscriptions table:", err)
    }
}