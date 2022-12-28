package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func GetConnection(dbName string) *sql.DB {
	connectionString := fmt.Sprintf("postgresql://postgres:postgres@localhost/%s?sslmode=disable", dbName)
	connection, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return connection
}
