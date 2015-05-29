package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())
	m.Get("/", func(r render.Render) {
		r.JSON(200, "Hello, Aaron Liang")
	})

	m.Get("/500", func(r render.Render) {
		r.Error(500)
	})

	m.Get("/502", func(r render.Render) {
		r.Error(502)
	})

	m.Get("/503", func(r render.Render) {
		r.Error(503)
	})

	m.Get("/504", func(r render.Render) {
		r.Error(504)
	})

	m.Run()
}
