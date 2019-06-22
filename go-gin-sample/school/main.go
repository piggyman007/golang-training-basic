package main

import (
	"database/sql"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Todo item
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func handleError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func getTodosHandler(c *gin.Context) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		handleError(c, err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos;")
	if err != nil {
		handleError(c, err)
		return
	}

	rows, err := stmt.Query()
	if err != nil {
		handleError(c, err)
		return
	}

	todos := make([]Todo, 0)
	for rows.Next() {
		todo := new(Todo)
		if err = rows.Scan(&todo.ID, &todo.Title, &todo.Status); err != nil {
			handleError(c, err)
			return
		}

		todos = append(todos, *todo)
	}

	c.JSON(http.StatusOK, todos)
}

func getTodosByIDHandler(c *gin.Context) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		handleError(c, err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos WHERE id = $1;")
	if err != nil {
		handleError(c, err)
		return
	}

	row := stmt.QueryRow(c.Param("id"))

	todo := new(Todo)
	if err = row.Scan(&todo.ID, &todo.Title, &todo.Status); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func postTodosHandler(c *gin.Context) {
	todo := new(Todo)

	if err := c.ShouldBindJSON(todo); err != nil {
		handleError(c, err)
		return
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		handleError(c, err)
	}
	defer db.Close()

	query := `
		INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id
	`
	row := db.QueryRow(query, todo.Title, todo.Status)
	if err = row.Scan(&todo.ID); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusAccepted, todo)
}

func updateTodosByIDHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handleError(c, err)
		return
	}

	todo := new(Todo)

	if err := c.ShouldBindJSON(todo); err != nil {
		handleError(c, err)
		return
	}
	todo.ID = id

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		handleError(c, err)
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
		handleError(c, err)
		return
	}

	if _, err = stmt.Exec(todo.ID, todo.Status, todo.Title); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func deleteTodosByIDHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handleError(c, err)
		return
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		handleError(c, err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`
		DELETE FROM todos WHERE id=$1;
	`)
	if err != nil {
		handleError(c, err)
		return
	}

	if _, err = stmt.Exec(id); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/todos", postTodosHandler)
		api.GET("/todos/:id", getTodosByIDHandler)
		api.GET("/todos", getTodosHandler)
		api.PUT("/todos/:id", updateTodosByIDHandler)
		api.DELETE("/todos/:id", deleteTodosByIDHandler)
	}

	r.Run(":1234")
}
