package main

import (
	"bufio"
	"os"
	"sync"
)

type game struct {
	buffer [][]byte
	player *player
}

func (g *game) ToString() string {
	return "asdf"
}

type player struct {
	x    int
	y    int
	char rune
}

var reader *bufio.Reader

var updateChannel = make(chan string, 20)

var group sync.WaitGroup

func main() {
	game := game{}
	p := player{
		x:    1,
		y:    1,
		char: 'X',
	}
	game.player = &p

	reader = bufio.NewReader(os.Stdin)

	game.buffer = loadMap()

	group.Add(1)
	renderMap(&game)
	go renderer(&game)
	go inputReader(&p)

	group.Wait()
}

func inputReader(p *player) {
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}

		if r == 'w' {
			p.y++
			if p.y > HEIGHT-1 {
				p.y = HEIGHT - 1
			}
		} else if r == 'a' {
			p.x--
			if p.x < 0 {
				p.x = 0
			}
		} else if r == 's' {
			p.y--
			if p.y < 0 {
				p.y = 0
			}
		} else if r == 'd' {
			p.x++
			if p.x > WIDTH-1 {
				p.x = WIDTH - 1
			}
		}

		updateChannel <- "yeah"
	}
}

func renderer(game *game) {
	for {
		<-updateChannel
		renderMap(game)
	}
}
