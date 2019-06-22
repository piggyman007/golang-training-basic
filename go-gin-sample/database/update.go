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
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("fatal ", err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE todos SET status=$2 WHERE id=$1;")
	if err != nil {
		log.Fatal("prepare error ", err.Error())
	}

	if _, err = stmt.Exec(6, "active"); err != nil {
		log.Fatal("exec error ", err.Error())
	}

	fmt.Println("updte success")
}
