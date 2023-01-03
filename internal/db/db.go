package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Connection *sql.DB

func EstablishConnection(databaseName string) {
	connectionString := fmt.Sprintf("postgresql://postgres:postgres@localhost/%s?sslmode=disable", databaseName)
	connection, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	Connection = connection
}
