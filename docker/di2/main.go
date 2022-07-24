package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// docker build -t web .
// docker run -p 8000:8080 web
func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello there!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
