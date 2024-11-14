package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func createNodeStructure(file *os.File) *Board {
	scanner := bufio.NewScanner(file)

	board := &Board{}

	var nodeList []*Node
	mapaCoors := make(map[string]*Node)
	isFirstLine := true

	for scanner.Scan() {
		line := scanner.Text()
		if isFirstLine {
			parts := strings.Split(line, "x")
			x, _ := strconv.ParseInt(parts[0], 10, 32)
			y, _ := strconv.ParseInt(parts[1], 10, 32)
			board.WIDTH = int(x)
			board.HEIGHT = int(y)
			isFirstLine = false
			continue
		}
		nodeTexts := strings.Split(line, " ")

		for _, node := range nodeTexts {
			parts := strings.Split(node, "-")
			x, _ := strconv.ParseInt(parts[0], 10, 32)
			y, _ := strconv.ParseInt(parts[1], 10, 32)
			isWalkable, _ := strconv.ParseInt(parts[2], 10, 32)

			newNode := &Node{
				x:        int(x),
				y:        int(y),
				cost:     1, //TODO hardcoded cost for the momment
				walkable: isWalkable == 1,
				edges:    make(map[*Node]int),
			}

			nodeList = append(nodeList, newNode)
			mapaCoors[newNode.getCoords()] = newNode
		}

	}

	board.coorsMap = mapaCoors
	for _, node := range nodeList {
		nList := node.neighborCoorKeys(board)
		for _, n := range nList {
			node.edges[n] = n.cost
		}
	}

	return board
}
