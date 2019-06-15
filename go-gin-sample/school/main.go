package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// type Student struct {
// 	Name string `json:"name" binding:"required"`
// 	ID   int    `json:"id" binding:"required"`
// }

var todos = map[int]Todo{}

// func getStudentsHandler(c *gin.Context) {
// 	ss := make([]Student, 0)
// 	for _, s := range students {
// 		ss = append(ss, s)
// 	}
// 	c.JSON(http.StatusOK, ss)
// }

// func postStudentsHandler(c *gin.Context) {
// 	s := new(Student)
// 	if err := c.ShouldBindJSON(s); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	id := len(students) + 1
// 	s.ID = id
// 	students[id] = *s
// 	c.JSON(http.StatusOK, s)
// }

func postTodosHandler(c *gin.Context) {
	todo := new(Todo)

	if err := c.ShouldBindJSON(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.ID = 1
	todos[todo.ID] = *todo
	c.JSON(http.StatusAccepted, todos[todo.ID])
}

func getTodosByIDHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos[id])
}

func getTodosHandler(c *gin.Context) {
	ts := make([]Todo, 0)

	for _, todo := range todos {
		ts = append(ts, todo)
	}

	c.JSON(http.StatusOK, ts)
}

func updateTodosByIDHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, ok := todos[id]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ID %d does not exists", id)})
		return
	}

	todo := new(Todo)
	todo.ID = id
	if err := c.ShouldBindJSON(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todos[id] = *todo
	c.JSON(http.StatusOK, todos[id])
}

func deleteTodosByIDHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, ok := todos[id]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ID %d does not exists", id)})
		return
	}

	delete(todos, id)
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
