package linkedlist

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
	if ll.size == 0 {
		ll.tail = newNode
	} else {
		ll.head.prev = newNode
	}
	ll.head = newNode
	ll.size++
}

func (ll *LinkedList) AddLast(value interface{}) {
	newNode := &node{value: value}
	newNode.prev = ll.tail
	if ll.size == 0 {
		ll.head = newNode
	} else {
		ll.tail.next = newNode
	}
	ll.tail = newNode
	ll.size++
}

func (ll *LinkedList) Size() int {
	return ll.size
}
