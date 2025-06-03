package main

import (
	"io"
	"text/template"

	"github.com/Maxi-Di-Mito/go-cli-game/entities"
	"github.com/Maxi-Di-Mito/go-cli-game/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var nodeList []*entities.Node

type Template struct {
	Templates *template.Template
}

var Temps = &Template{
	Templates: template.Must(template.ParseGlob("views/*.go.html")),
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

var theGame *entities.Game

func main() {
	server := echo.New()

	server.Use(middleware.Logger())
	server.Static("/static", "static")
	server.Renderer = Temps

	server.GET("/", routes.HomeHandler)

	server.GET("/api/getMap", routes.GetMapHandlerMaker(&theGame))

	server.POST("/api/click", routes.ClickHandlerMaker(&theGame))

	// theGame := startGame()
	// path := theGame.getPath("0-0", "4-3")
	// for _, node := range path {
	// 	fmt.Printf("%d-%d\n", node.x, node.y)
	// }

	server.Start(":1234")

}
