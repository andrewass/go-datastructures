package graph

import (
	"go-datastructures-algorithms/list/arraylist"
	"go-datastructures-algorithms/stack"
)

type node struct {
	item     interface{}
	key      interface{}
	outEdges *arraylist.ArrayList
	distance *int64
	parent   *node
	visited  bool
	tempVisited bool
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

// AddItem inserts a new item in the graph, wrapped in a new node
func (g *Graph) AddItem(key interface{}, item interface{}) {
	node := &node{item: item, key: key, outEdges: arraylist.New()}
	g.nodes[key] = node
	g.size++
}

// AddEdge adds an edge between two items in the graph
func (g *Graph) AddEdge(fromKey interface{}, toKey interface{}, weight int64) {
	fromNode := g.nodes[fromKey]
	toNode := g.nodes[toKey]
	edge := &edge{from: fromNode, to: toNode, weight: &weight}
	fromNode.outEdges.Add(edge)
}

// GetShortestDistance returns the shortest distance between two items in the graph.
// If no path exists between source and destination, nil is returned
func (g *Graph) GetShortestDistance(sourceKey, destinationKey interface{}) *int64 {
	sourceNode := g.nodes[sourceKey]
	endNode := g.nodes[destinationKey]
	setShortestPath(sourceNode, g)

	if *endNode.distance == *getMaxDistance(){
		return nil
	}
	return endNode.distance
}

// GetShortestDistanceToAll returns a list of nodes, each holding its shortest distance from the given source
func (g *Graph) GetShortestDistanceToAll(source interface{}) *arraylist.ArrayList {
	return nil
}

func (g *Graph) GetTopologicalOrder() *arraylist.ArrayList {
	return topologicalSort(g)
}

// ExistsPath return true if there exists a path from source to end node
func (g *Graph) ExistsPath(fromKey, toKey interface{}) bool {
	sourceNode := g.nodes[fromKey]
	nodeStack := stack.New()
	nodeStack.Push(sourceNode)
	g.clearVisitedFlags()

	for !nodeStack.IsEmpty() {
		current := nodeStack.Pop().(*node)
		if current.key == toKey {
			return true
		}
		if !current.visited {
			current.visited = true
			for i := 0; i < current.outEdges.Size(); i++ {
				edge := current.outEdges.Get(i).(*edge)
				nodeStack.Push(edge.to)
			}
		}
	}
	return false
}

func (g *Graph) clearVisitedFlags() {
	for _, node := range g.nodes {
		node.visited = false
	}
}

func getMaxDistance() *int64 {
	var maxValue int64 = 100000000000
	return &maxValue
}