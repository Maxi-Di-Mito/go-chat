package main

import (
	"database/sql"
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type Price struct {
	Price int
	Name  string
}

var DB *sql.DB

type Template struct {
	Templates *template.Template
}

var Temps = &Template{
	Templates: template.Must(template.ParseGlob("templates/*.go.html")),
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func HomeHandler(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "index", nil)
}

func CreateHandler(ctx echo.Context) error {
	if ctx.Request().Method == "POST" {
		name := ctx.FormValue("name")                                                // Get the title from the form data
		price := ctx.FormValue("price")                                              // Get the title from the form data
		_, err := DB.Exec("INSERT INTO prices(name,price) VALUES(?,?)", name, price) // Insert the new todo into the database
		if err != nil {
			return err
		}
		return ctx.Render(http.StatusOK, "index", nil) // Redirect to the main page after successful creation
	}
	return ctx.NoContent(http.StatusBadRequest)
}

func main() {
	InitDB()
	defer DB.Close()

	server := echo.New()
	server.Use(middleware.Logger())
	server.Static("/static", "static")
	server.Renderer = Temps

	server.GET("/", HomeHandler)
	server.POST("/create", HomeHandler)

	server.Start(":1234")
}
