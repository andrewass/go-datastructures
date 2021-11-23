package graph

import (
	"go-datastructures-algorithms/list/arraylist"
)

type node struct {
	item     interface{}
	outEdges *arraylist.ArrayList
	distance *int64
	parent   *node
}

type edge struct {
	from   *node
	to     *node
	weight *int64
}

type Graph struct {
	nodes map[interface{}]*node
	size  int
}

func New() *Graph {
	return &Graph{nodes: make(map[interface{}]*node)}
}

func (g *Graph) AddItem(key interface{}, item interface{}) {
	node := &node{item: item, outEdges: arraylist.New()}
	g.nodes[key] = node
	g.size++
}

// AddEdge adds an edge between two items in the graph
func (g *Graph) AddEdge(from interface{}, to interface{}, weight int64) {
	fromNode := g.nodes[from]
	toNode := g.nodes[to]
	edge := &edge{from: fromNode, to: toNode, weight: &weight}
	toNode.outEdges.Add(edge)
}

// GetDistance returns the shortest distance between two items in the graph.
func (g *Graph) GetDistance(source, end interface{}) *int64 {
	return nil
}

// GetDistanceToAll returns a list of nodes, each holding its shortest distance from the given source
func (g *Graph) GetDistanceToAll(source interface{}) *arraylist.ArrayList {
	return nil
}
