package entities

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Board struct {
	coorsMap map[string]*Node
	WIDTH    int
	HEIGHT   int
}

func (board *Board) ToArray() [][]string {
	ret := make([][]string, board.WIDTH)
	for idx, row := range ret {
		row = make([]string, board.HEIGHT)
		ret[idx] = row
	}

	for _, n := range board.coorsMap {
		if n.walkable {
			ret[n.x][n.y] = "1"
		} else {
			ret[n.x][n.y] = "0"
		}
	}

	return ret
}

func (board *Board) Dijkstra(startKey string) (result *DijkstraResult, err error) {
	initial, ok := board.coorsMap[startKey]
	if !ok {
		return nil, fmt.Errorf("start vertex %v not found", startKey)
	}

	distances := make(map[string]int)
	prevs := make(map[*Node]*Node)
	prevs[initial] = nil
	for key := range board.coorsMap {
		if board.coorsMap[key].walkable { // avoid assigning infinite to unwalkable nodes
			distances[key] = math.MaxInt32
		}
	}
	distances[startKey] = 0

	var vertices []*Node
	for _, vertex := range board.coorsMap {
		if vertex.walkable { // avoid calculating distance to unwalkable nodes
			vertices = append(vertices, vertex)
		}
	}

	for len(vertices) != 0 {
		sort.SliceStable(vertices, func(i, j int) bool {
			return distances[vertices[i].getKey()] < distances[vertices[j].getKey()]
		})

		vertex := vertices[0]
		vertices = vertices[1:]

		for adjacent, cost := range vertex.edges {
			alt := distances[vertex.getKey()] + cost
			if alt < distances[adjacent.getKey()] {
				distances[adjacent.getKey()] = alt
				prevs[adjacent] = vertex
			}
		}
	}

	result = &DijkstraResult{
		distances: distances, prevs: prevs}

	return result, err
}

func (board *Board) getPath(key string, result *DijkstraResult) []*Node {
	path := []*Node{}
	target, ok := board.coorsMap[key]
	if !ok {
		panic("No target not in board")
	}
	path = append(path, target)
	for prev, ok := result.prevs[target]; ok && prev != nil; {
		path = append(path, prev)
		// fmt.Println("PREV", prev)
		prev, ok = result.prevs[prev]
	}
	slices.Reverse(path)

	return path
}

type DijkstraResult struct {
	distances map[string]int
	prevs     map[*Node]*Node
}

func createNodeStructure(file *os.File) *Board {
	scanner := bufio.NewScanner(file)

	board := &Board{}

	var nodeList []*Node
	mapaCoors := make(map[string]*Node)

	for yCoor := -1; scanner.Scan(); yCoor++ {
		line := scanner.Text()
		if yCoor == -1 {
			parts := strings.Split(line, "x")
			x, _ := strconv.ParseInt(parts[0], 10, 32)
			y, _ := strconv.ParseInt(parts[1], 10, 32)
			board.WIDTH = int(x)
			board.HEIGHT = int(y)
			continue
		}
		nodeTexts := strings.Split(line, " ")

		for xCoor, node := range nodeTexts {
			isWalkable, _ := strconv.ParseInt(node, 10, 32)

			newNode := &Node{
				x:        xCoor,
				y:        yCoor,
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
