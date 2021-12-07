package graph

import (
	"go-datastructures-algorithms/list/arraylist"
)

func topologicalSort(graph *Graph) *arraylist.ArrayList {
	sortedList := arraylist.New()
	initializeGraphTopologicalSort(graph)

	for _, node := range graph.nodes {
		if !node.visited {
			visitNode(node, sortedList)
		}
	}
	sortedList.Reverse()
	return sortedList
}

func visitNode(current *node, sortedList *arraylist.ArrayList) {
	if current.visited {
		return
	}
	if current.tempVisited {
		panic("Unable to sort a graph with a cycle")
	}
	current.tempVisited = true
	for i := 0; i < current.outEdges.Size(); i++ {
		edge := current.outEdges.Get(i).(*edge)
		visitNode(edge.to, sortedList)
	}
	current.tempVisited = false
	current.visited = true
	sortedList.Add(current.item)
}

func initializeGraphTopologicalSort(graph *Graph) {
	for _, node := range graph.nodes {
		node.visited = false
		node.tempVisited = false
	}
}
