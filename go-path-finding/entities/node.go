package entities

import (
	"fmt"
	"slices"
)

type Node struct {
	x        int
	y        int
	cost     int
	edges    map[*Node]int
	walkable bool
	state    string
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
	if n.x > 0 && board.coorsMap[n.getLeftCoors()].walkable {
		neighbors = append(neighbors, board.coorsMap[n.getLeftCoors()])
	}

	neighbors = slices.DeleteFunc(neighbors, func(n *Node) bool {
		return !n.walkable
	})

	return neighbors
}
