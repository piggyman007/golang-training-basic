package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/piggyman007/school/database"
)

func GetList(c *gin.Context) {
	rows, err := db.GetTodoList()
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
