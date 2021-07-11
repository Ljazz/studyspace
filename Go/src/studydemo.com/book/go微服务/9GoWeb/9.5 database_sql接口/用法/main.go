package main

import (
	// Setp 1: import the main SQL package
	"database/sql"
	// Setp 2: import a driver package to use a specific SQL databases
	_ "github.com/mattn/go-sqlite3"
)

// Setp 3: open a database using a registered driver name

func main() {
	// ...
	db, err := sql.Open("sqlite3", "database.db")
	// ...
}
