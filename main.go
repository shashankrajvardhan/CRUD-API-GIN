package main

import "database/sql"

var db *sql.DB
var students []Students

func main() {
	connection()
	router()
}
