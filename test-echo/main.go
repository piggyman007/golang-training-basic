package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	md := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// fmt.Println(">>>> middleware")
			// return next(c)
			// return c.String(http.StatusUnauthorized, "Failed!")
			return echo.NewHTTPError(http.StatusBadRequest, "bad request")

		}
	}

	g := e.Group("/api")
	g.GET("/test", func(c echo.Context) error {
		fmt.Println(">>>> handler")
		return c.String(http.StatusOK, "Hello, World!")
	}, md)

	e.Logger.Fatal(e.Start(":1323"))
}
