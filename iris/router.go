package main

import (
	"github.com/kataras/iris"
)

func main() {

	iris.Get("/users", func(c *iris.Context) {
		c.Write("All users")
	})

	iris.Get("/users/*id", func(c *iris.Context) {
		c.Write("User %s", c.Param("id"))
	})

	iris.Get("/users/aaron", func(c *iris.Context) {
		c.Write("Hello aaron")
	})

	iris.Listen(":8080")
}
