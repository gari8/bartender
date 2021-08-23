package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDatabaseConnection() (*sql.DB, error) {
	conn, err := sql.Open("postgres", "host=db user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func main() {
	conn, _ := NewDatabaseConnection()

}
