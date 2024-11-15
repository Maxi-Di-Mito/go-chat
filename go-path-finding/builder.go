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

	for x := -1; scanner.Scan(); x++ {
		line := scanner.Text()
		if x == -1 {
			parts := strings.Split(line, "x")
			x, _ := strconv.ParseInt(parts[0], 10, 32)
			y, _ := strconv.ParseInt(parts[1], 10, 32)
			board.WIDTH = int(x)
			board.HEIGHT = int(y)
			continue
		}
		nodeTexts := strings.Split(line, " ")

		for y, node := range nodeTexts {
			isWalkable, _ := strconv.ParseInt(node, 10, 32)

			newNode := &Node{
				x:        x,
				y:        y,
				cost:     1, //TODO hardcoded cost for the momment
				walkable: isWalkable == 1,
				edges:    make(map[*Node]int),
			}

			nodeList = append(nodeList, newNode)
			mapaCoors[newNode.getKey()] = newNode
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
