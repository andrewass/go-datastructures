package graph

import (
	"go-datastructures-algorithms/queue/priorityqueue"
)

func distanceComparator(a, b interface{}) int {
	first := a.(*node)
	second := b.(*node)

	if *first.distance < *second.distance {
		return -1
	}
	return 0
}

func setShortestPath(source *node, graph *Graph) {
	priQueue := priorityqueue.New(distanceComparator)
	initializeGraphShortestPath(source, graph)
	priQueue.Insert(source)

	for !priQueue.IsEmpty() {
		current := priQueue.PollMin().(*node)
		for i := 0; i < current.outEdges.Size(); i++ {
			edge := current.outEdges.Get(i).(*edge)
			distance := *current.distance + *edge.weight
			if distance < *edge.to.distance {
				edge.to.distance = &distance
				edge.to.parent = current
				priQueue.Insert(edge.to)
			}
		}
	}
}

func initializeGraphShortestPath(source *node, graph *Graph) {
	var zeroDistance int64 = 0
	for _, node := range graph.nodes {
		node.distance = getMaxDistance()
		if node == source {
			node.distance = &zeroDistance
		}
	}
}
