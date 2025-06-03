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

func GetMapHandlerMaker(game **entities.Game) func(echo.Context) error {
	return func(ctx echo.Context) error {
		if *game == nil {
			*game = entities.StartGame()
		}

		ctx.Response().Header().Set("Content-Type", "application/json")

		return ctx.JSON(http.StatusOK, (*game).Board.ToArray())
	}
}

func ClickHandlerMaker(game **entities.Game) func(echo.Context) error {
	return func(ctx echo.Context) error {
		// access request body params
		var body struct {
			X int `json:"x"`
			Y int `json:"y"`
		}
		if err := ctx.Bind(&body); err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid request")
		}

		x := body.X
		y := body.Y

		coor := fmt.Sprintf("%d-%d", x, y)
		if game == nil {
			fmt.Println("Game is nil")
		}

		var response struct {
			Next string `json:"next"`
			Coor string `json:"coor"`
		}
		response.Coor = coor

		if (*game).Start != "" {
			(*game).Target = coor
			(*game).GetPath((*game).Start, (*game).Target)
			response.Next = "path"
			return ctx.JSON(http.StatusOK, response)
		} else {
			(*game).Start = coor
			response.Next = "target"
			return ctx.JSON(http.StatusOK, response)
		}

	}
}
