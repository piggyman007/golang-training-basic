package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Todo struct {
	id     int
	title  string
	status string
}

func main() {
	// url := "postgres://cbrowfvv:7jKJjXxbXvUaD8hNvIwYdiUJXLh0Kgq4@satao.db.elephantsql.com:5432/cbrowfvv"
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("fatal ", err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos WHERE id = $1;")
	if err != nil {
		log.Fatal("prepare error ", err.Error())
	}

	row := stmt.QueryRow(6)

	t := new(Todo)

	if err = row.Scan(&t.id, &t.title, &t.status); err != nil {
		log.Fatal("error ", err.Error())
	}

	fmt.Printf("one row => %+v\n", t)
}
