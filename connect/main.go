package main

import "database/sql"
import _ "github.com/trinodb/trino-go-client/trino"

func main() {
	dsn := "http://user@localhost:8080?catalog=default&schema=test"
	db, err := sql.Open("trino", dsn)
}
