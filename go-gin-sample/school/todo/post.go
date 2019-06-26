package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/piggyman007/school/database"
)

func Post(c *gin.Context) {
	todo := new(Todo)

	if err := c.ShouldBindJSON(todo); err != nil {
		handleError(c, err)
		return
	}

	row, err := db.CreateTodo(todo.Title, todo.Status)

	if err != nil {
		handleError(c, err)
		return
	}

	if err = row.Scan(&todo.ID); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusAccepted, todo)
}
