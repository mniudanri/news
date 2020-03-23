package connection

import (
	"fmt"
	"database/sql"
    // "html/template"
    _ "github.com/lib/pq"
)
func Connect() (*sql.DB) {
	var err error
    db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres "+
            "password=12345678 dbname=masterdata sslmode=disable")

    if err != nil {
        panic(err)
    }

    if err = db.Ping(); err != nil {
        panic(err)
    }

    fmt.Println("Succesfully connected to database.")

	return db
}