package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// url := "postgres://cbrowfvv:7jKJjXxbXvUaD8hNvIwYdiUJXLh0Kgq4@satao.db.elephantsql.com:5432/cbrowfvv"
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("fatal ", err.Error())
	}
	defer db.Close()

	title := "Home Work"
	status := "Inactive"
	query := `
		INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id
	`
	var id int
	row := db.QueryRow(query, title, status)
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("can't scan id", id)
	}

	fmt.Println("insert success id: ", id)
}
