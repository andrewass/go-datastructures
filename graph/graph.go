package graph

import "go-datastructures-algorithms/list/arraylist"

type Node struct {
	item     interface{}
	outEdges *arraylist.ArrayList
}

type Edge struct {
	from   *Node
	to     *Node
	length int
}

type Graph struct {
	nodes map[interface{}]*Node
	size  int
}

func New() *Graph {
	return &Graph{nodes: make(map[interface{}]*Node)}
}

func (g *Graph) AddNode(key interface{}, item interface{}) {
	node := &Node{item: item, outEdges: arraylist.New()}
	g.nodes[key] = node
	g.size++
}

func (g *Graph) AddEdge(from interface{}, to interface{}, length int) {
	fromNode := g.nodes[from]
	toNode := g.nodes[to]
	edge := &Edge{from: fromNode, to: toNode, length: length}
	toNode.outEdges.Add(edge)
}
