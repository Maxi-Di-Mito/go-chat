package entities

import "os"

type Game struct {
	Board     *Board
	Player    []int
	Dijkstras map[string]*DijkstraResult
	Start     string
	Target    string
}

func (g *Game) InitMap() {
	g.Board = LoadMap()
}

func (g *Game) GetPath(from string, to string) []*Node {
	if _, ok := g.Dijkstras[from]; !ok {
		newResult, _ := g.Board.Dijkstra(from)
		g.Dijkstras[from] = newResult
	}
	path := g.Board.getPath(to, g.Dijkstras[from])

	return path
}

func StartGame() *Game {
	game := Game{}
	game.InitMap()
	game.Dijkstras = make(map[string]*DijkstraResult)

	return &game
}

func LoadMap() *Board {
	file, err := os.Open("map.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	board := createNodeStructure(file)

	return board
}
