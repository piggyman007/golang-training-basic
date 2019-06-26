package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/piggyman007/school/database"
)

func DeleteByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handleError(c, err)
		return
	}

	if err = db.DeleteTodoByID(id); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
