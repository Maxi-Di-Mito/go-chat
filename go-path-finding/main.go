package main

import (
	"fmt"
	"os"
)

var nodeList []*Node

func main() {
	file, err := os.Open("map.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	board := createNodeStructure(file)

	fmt.Printf("%+v\n", board.coorsMap)

	result, err := board.Dijkstra("0-0")
	if err != nil {
		panic(err)
	}

	fmt.Printf("DISTANCES:\n%+v\n", result.distances)

	fmt.Println("Path to 2-2")
	path := board.getPath("2-2", result)

	for _, p := range path {
		fmt.Printf("%+v - ", p.getKey())
	}

}
