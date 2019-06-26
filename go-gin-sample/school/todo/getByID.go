package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/piggyman007/school/database"
)

func GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handleError(c, err)
		return
	}

	row, err := db.GetTodoByID(id)
	if err != nil {
		handleError(c, err)
		return
	}

	todo := new(Todo)
	if err = row.Scan(&todo.ID, &todo.Title, &todo.Status); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}
