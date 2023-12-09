package lib

import "fmt"

type Node struct {
	Name  string // Name of the node
	Left  *Node
	Right *Node
}

// Graph represents network of nodes.
type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

// AddNode adds a new node to the graph.
func (g *Graph) AddNode(name string) *Node {
	node := &Node{Name: name}
	g.Nodes[name] = node
	return node
}

// AddEdge creates an edge between two nodes.
func (g *Graph) AddEdge(from, to string, left bool) {
	fromNode, fromExists := g.Nodes[from]
	toNode, toExists := g.Nodes[to]

	if !fromExists || !toExists {
		fmt.Printf("Warning: Trying to add edge from %s to %s, but one of the nodes does not exist.\n", from, to)
		return
	}

	if left {
		fromNode.Left = toNode
	} else {
		fromNode.Right = toNode
	}
}
