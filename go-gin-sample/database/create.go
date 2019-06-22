package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	url := "postgres://cbrowfvv:7jKJjXxbXvUaD8hNvIwYdiUJXLh0Kgq4@satao.db.elephantsql.com:5432/cbrowfvv"
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("fatal ", err.Error())
	}
	defer db.Close()

	createTb := `
		CREATE TABLE IF NOT EXISTS todos(
			id		SERIAL PRIMARY KEY,
			title	TEXT,
			status	TEXT
		);
	`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table ", err.Error())
	}

	fmt.Println("OK")
}
