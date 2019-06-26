package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/piggyman007/school/database"
)

func UpdateByID(c *gin.Context) {
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

	err = db.UpdateTodoByID(todo.ID, todo.Status, todo.Title)

	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}
