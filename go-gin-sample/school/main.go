package main

import (
	"fmt"
	"net/http"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/piggyman007/school/todo"
)

func main() {
	hs256 := jwt.NewHMAC(jwt.SHA256, []byte("secret"))
	h := jwt.Header{}
	p := struct {
		id int
	}{
		id: 1,
	}
	token, _ := jwt.Sign(h, p, hs256)
	fmt.Printf("token = %s", token)

	raw, _ := jwt.Parse(token)
	if err := raw.Verify(hs256); err != nil {
		fmt.Println("verify failed")
	} else {
		fmt.Println("verify success")
	}

	r := setupRouter()
	r.Run(":1234")
}

func authMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "Bearer token123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		c.Abort()
		return
	}

	c.Next()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(authMiddleware)

	api := r.Group("/api")
	{
		api.POST("/todos", todo.Post)
		api.GET("/todos/:id", todo.GetByID)
		api.GET("/todos", todo.GetList)
		api.PUT("/todos/:id", todo.UpdateByID)
		api.DELETE("/todos/:id", todo.DeleteByID)
		api.GET("/ping", func(c *gin.Context) {
			fmt.Println(">> handler")
			c.JSON(http.StatusOK, gin.H{
				"hello": "world",
			})
		})
	}

	return r
}
