package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	Templates *template.Template
}

var Temps = &Template{
	Templates: template.Must(template.ParseGlob("views/*.go.html")),
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func main() {
	server := echo.New()
	server.Use(middleware.Logger())
	server.Static("/static", "static")
	server.Renderer = Temps
	server.GET("/", handleHome)
}

func handleHome(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}
