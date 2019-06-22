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

	stmt, err := db.Prepare("SELECT id, title, status FROM todos;")
	if err != nil {
		log.Fatal("prepare error ", err.Error())
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("query error ", err.Error())
	}

	tds := make([]Todo, 0)
	for rows.Next() {
		td := new(Todo)
		if err = rows.Scan(&td.id, &td.title, &td.status); err != nil {
			log.Fatal("scan error ", err.Error())
		}

		tds = append(tds, *td)
	}

	fmt.Printf("% +v\n", tds)
}
