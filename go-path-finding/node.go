package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

type Node struct {
	key      string
	x        int
	y        int
	cost     int
	edges    map[*Node]int
	walkable bool
}

func (n *Node) getKey() string {
	return fmt.Sprintf("%d-%d", n.x, n.y)
}

func (n *Node) getUpCoors() string {
	return fmt.Sprintf("%d-%d", n.x, n.y-1)
}
func (n *Node) getRightCoors() string {
	return fmt.Sprintf("%d-%d", n.x+1, n.y)
}
func (n *Node) getDownCoors() string {
	return fmt.Sprintf("%d-%d", n.x, n.y+1)
}
func (n *Node) getLeftCoors() string {
	return fmt.Sprintf("%d-%d", n.x-1, n.y)
}

func (n *Node) neighborCoorKeys(board *Board) []*Node {
	neighbors := []*Node{}
	if n.y > 0 && board.coorsMap[n.getUpCoors()].walkable {
		neighbors = append(neighbors, board.coorsMap[n.getUpCoors()])
	}
	if n.x < board.WIDTH-1 && board.coorsMap[n.getRightCoors()].walkable {
		neighbors = append(neighbors, board.coorsMap[n.getRightCoors()])
	}
	if n.y < board.HEIGHT-1 && board.coorsMap[n.getDownCoors()].walkable {
		neighbors = append(neighbors, board.coorsMap[n.getDownCoors()])
	}
	fmt.Printf("%+v\n", *n)
	fmt.Printf("LEFT COORDS %s", n.getLeftCoors())
	if n.x > 0 && board.coorsMap[n.getLeftCoors()].walkable {
		neighbors = append(neighbors, board.coorsMap[n.getLeftCoors()])
	}

	neighbors = slices.DeleteFunc(neighbors, func(n *Node) bool {
		return !n.walkable
	})

	return neighbors
}

type Board struct {
	coorsMap map[string]*Node
	WIDTH    int
	HEIGHT   int
}

type DjistraResult struct {
	distances map[string]int
	prevs     map[*Node]*Node
}

func (board *Board) Dijkstra(startKey string) (result *DjistraResult, err error) {
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

	result = &DjistraResult{
		distances: distances, prevs: prevs}

	return result, err
}

func (board *Board) getPath(key string, result *DjistraResult) []*Node {
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
