package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	// "errors"
)

var db *sql.DB

var server = "localhost"
var port = 1433

// var database = "Services"

func main() {
	var err error

	// Build connection string
	connString := fmt.Sprintf("server=%s;port=%d;trusted_connection=yes", server, port)

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	fmt.Printf("Connected!\n")

	// Close the database connection pool after program executes
	defer db.Close()
}
