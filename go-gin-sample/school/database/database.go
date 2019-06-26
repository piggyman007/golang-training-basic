package database

import (
	"database/sql"
	"os"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	return db, err
}

func UpdateTodoByID(id int, status string, title string) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(`
		UPDATE todos 
		SET 
			status=$2,
			title=$3
		WHERE id=$1;
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, status, title)

	return err
}

func CreateTodo(title string, status string) (*sql.Row, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
		INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id
	`
	row := db.QueryRow(query, title, status)

	return row, nil
}

func GetTodoByID(id int) (*sql.Row, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos WHERE id = $1;")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)

	return row, nil
}

func DeleteTodoByID(id int) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(`
		DELETE FROM todos WHERE id=$1;
	`)
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func GetTodoList() (*sql.Rows, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos;")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	return rows, err
}
