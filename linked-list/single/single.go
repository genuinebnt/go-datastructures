package singleLinkedList

//List contain pointer to beginning and end of list
type List struct {
	head *Node //head points to start of list
	tail *Node // tail points to end of list
}

//Node contains a pointer to another node or nil and a element
type Node struct {
	elem int
	next *Node //points to next node
}

//InsertEnd insert element at end of list
func (l *List) InsertEnd(elem int) {
	node := Node{elem: elem, next: nil} //initialize node with input

	if l.head == nil {
		l.head = &node //make head point to first node created
	}

	if l.tail == nil {
		l.tail = &node //make tail point to next node created
	}

	l.tail.next = &node  //make the node that tail points to, point to new node. so node is inserted at the end
	l.tail = l.tail.next //tail points to new node
}

//Show lists all elements inside the linked list
func (l *List) Show() []int {
	currentNode := l.head

	if currentNode == nil {
		return []int{}
	}

	result := []int{}
	for {
		result = append(result, currentNode.elem)
		currentNode = currentNode.next
		if currentNode.next == nil {
			result = append(result, currentNode.elem)
			break
		}
	}

	return result

}
