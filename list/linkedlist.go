package list

type LinkedList struct {
	head *node
	tail *node
	size int
}

type node struct {
	value interface{}
	next  *node
	prev *node
}

func New() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) PeekFirst() interface{} {
	if ll.size == 0{
		return nil
	}
	return ll.head.value
}

func (ll *LinkedList) PollFirst() interface{} {
	if ll.size == 0{
		return nil
	}
	headNode := ll.head
	ll.head = headNode.next
	ll.size--

	return headNode.value
}

func (ll *LinkedList) AddFirst(value interface{}) {
	newNode := &node{value: value}
	newNode.next = ll.head
	ll.head = newNode
	if ll.size == 0 {
		ll.tail = newNode
	}
	ll.size++
}

func (ll *LinkedList) AddLast(value interface{}) {

}

func (ll *LinkedList) Size() int {
	return ll.size
}
