package main

import (
	"strings"

	"github.com/Maxi-Di-Mito/go-cli-game/entities"
)

const BLANK = " "
const WALL = "W"
const Path = "o"

func Render(game *entities.Game, path []*entities.Node) string {
	ren := strings.Repeat(strings.Repeat(BLANK, game.Board.WIDTH), game.Board.HEIGHT)

	return ren
}
