package routes

import (
	"fmt"
	"net/http"

	"github.com/Maxi-Di-Mito/go-cli-game/entities"
	"github.com/labstack/echo"
)

func HomeHandler(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "index", nil)
}

func GetMapHandlerMaker(game *entities.Game) func(echo.Context) error {
	return func(ctx echo.Context) error {
		if game == nil {
			game = entities.StartGame()
		}

		ctx.Response().Header().Set("Content-Type", "application/json")

		return ctx.JSON(http.StatusOK, game.Board.ToArray())
	}
}

func ClickHandlerMaker(game *entities.Game) func(echo.Context) error {
	return func(ctx echo.Context) error {
		x := ctx.QueryParam("X")
		y := ctx.QueryParam("Y")

		coor := fmt.Sprintf("%s-%s", x, y)

		if game.Start != "" {
			game.Target = coor
			game.GetPath(game.Start, game.Target)
			return ctx.String(http.StatusOK, coor)
		} else {
			game.Start = coor
			return ctx.String(http.StatusOK, coor)
		}

	}
}
