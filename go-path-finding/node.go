package main

type Node struct {
	key   string
	x     int
	y     int
	board *Board
	cost  int
	edges map[*Node]int
}

func (n *Node) topNode() *Node {
	if n.y == 0 {
		return nil
	}
	return n.board.nodes[n.x][n.y-1]
}
func (n *Node) bottomNode() *Node {
	if n.y+1 == HEIGHT {
		return nil
	}
	return n.board.nodes[n.x][n.y+1]
}
func (n *Node) leftNode() *Node {
	if n.x == 0 {
		return nil
	}
	return n.board.nodes[n.x-1][n.y]
}
func (n *Node) rightNode() *Node {
	if n.x+1 == WIDTH {
		return nil
	}
	return n.board.nodes[n.x+1][n.y]
}
func (n *Node) isEnd() bool {
	return n.key == "END"
}

type Board struct {
	graph  *Graph
	WIDTH  int
	HEIGHT int
}

func (b *Board) toGraph() *Graph {

	return nil
}
