package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	WIDTH   = 20
	HEIGHT  = 20
	NOTHING = 0
	WALL    = 1
)

func loadMap() [][]byte {
	file, err := os.Open("map.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	buf := make([][]byte, HEIGHT)

	h := 0

	for scanner.Scan() {
		line := scanner.Text()
		buf[h] = []byte(line)
		h++
	}

	return buf
}

func renderMap(game *game) {

	mapa := game.buffer

	mapa[game.player.y][game.player.x] = byte(game.player.char)

	fmt.Printf("%+v\n", game)

	for r := 0; r < WIDTH; r++ {
		if r == game.player.y {
			fmt.Println("HEY")
		}
		row := string(game.buffer[r])
		fmt.Println(row)
	}
}
