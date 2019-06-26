package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/piggyman007/school/todo"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/todos", todo.Post)
		api.GET("/todos/:id", todo.GetByID)
		api.GET("/todos", todo.GetList)
		api.PUT("/todos/:id", todo.UpdateByID)
		api.DELETE("/todos/:id", todo.DeleteByID)
	}

	r.Run(":1234")
}
