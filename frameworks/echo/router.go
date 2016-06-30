package main

import (
	"net/http"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[Echo] ${time_rfc3339} | ${status} ${method} - ${uri} | ${latency_human} | ${tx_bytes} bytes\n",
	}))

	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, users")
	})

	e.GET("/users/aaron", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, aaron")
	})

	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, "Hello, user: "+id)
	})

	e.GET("/users/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, new user")
	})

	e.Run(standard.New(":8080"))
}
