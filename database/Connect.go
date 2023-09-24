package database

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

const (
    username = "bursaweb_ajans"
    password = "Genetik1997.*/"
    hostname = "84.54.13.3"
    port     = 3306
    database = "bursaweb_crm"
)

func OpenDB() (*sql.DB, error) {
    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, hostname, port, database)

    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    fmt.Println("MySQL veritabanına başarıyla bağlandınız.")
    return db, nil
}
