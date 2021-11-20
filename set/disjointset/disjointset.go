package disjointset

type setNode struct {
	item   interface{}
	rank   int
	size   int
	parent *setNode
}

type DisjointSet struct {
	nodeSet  []*setNode
	mapNodes map[interface{}]*setNode
}

func New() *DisjointSet {
	return &DisjointSet{mapNodes: make(map[interface{}]*setNode)}
}

func (djs *DisjointSet) AddItem(key interface{}, item interface{}) {
	newNode := &setNode{item: item, size: 1}
	djs.nodeSet = append(djs.nodeSet, newNode)
	djs.mapNodes[key] = newNode
}

func (djs *DisjointSet) GetItem(key interface{}) interface{} {
	node := djs.mapNodes[key]
	if node == nil {
		return node
	} else {
		return node.item
	}
}

func (djs *DisjointSet) Union(keyA interface{}, keyB interface{}) {
	rootA := djs.findRootNode(djs.mapNodes[keyA])
	rootB := djs.findRootNode(djs.mapNodes[keyB])

	if rootA == rootB {
		return
	}
	if rootA.rank > rootB.rank {
		rootB.parent = rootA
		rootA.size += rootB.size
	} else {
		rootA.parent = rootB
		rootB.size += rootA.size
		if rootB.rank == rootA.rank {
			rootB.rank++
		}
	}
}

func (djs *DisjointSet) IsSameSet(keyA interface{}, keyB interface{}) bool {
	return djs.findRootNode(djs.mapNodes[keyA]) == djs.findRootNode(djs.mapNodes[keyB])
}

func (djs *DisjointSet) GetSetSize(key interface{}) int {
	rootNode := djs.findRootNode(djs.mapNodes[key])

	return rootNode.size
}

func (djs *DisjointSet) findRootNode(node *setNode) *setNode {
	if node.parent != nil{
		node = djs.findRootNode(node.parent)
	}
	return node
}
