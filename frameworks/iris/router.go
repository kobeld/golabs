package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.Get("/users", func(c iris.Context) {
		c.WriteString("All users")
	})

	app.Get("/users/{id:uint}", func(c iris.Context) {
		userID := c.Params().GetUintDefault("id", 0)
		c.Writef("UserID: %d", userID)
	})

	app.Get("/users/aaron", func(c iris.Context) {
		c.WriteString("Hello aaron")
	})

	app.Run(iris.Addr(":8080"))
}
