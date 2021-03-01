package circularsinglylinkedlist

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

	if l.head == nil || l.tail == nil {
		l.head = &node //make head point to first node created
		l.tail = &node //make tail point to first node created
	} else {
		l.tail.next = &node  //make the node that tail points to, point to new node. so node is inserted at the end
		l.tail = l.tail.next //tail points to new node
		node.next = l.head   //last node point to first node
	}
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
		if currentNode.next == l.head { //if current node points to first node(circular), stop
			result = append(result, currentNode.elem)
			break
		}
	}

	return result

}

//InsertFront add node as first element of linked list
func (l *List) InsertFront(elem int) {

	node := Node{elem: elem, next: nil} // inititialize node with element

	if l.head == nil || l.tail == nil {
		l.head = &node
		l.tail = &node
	} else {
		node.next = l.head   //make new node point to node head points to
		l.head = &node       //make head point to new node
		l.tail.next = l.head //make last node point to first node
	}

}

//Search loop through list to find element
func (l *List) Search(elem int) string {
	currentNode := l.head
	if currentNode == nil {
		return "Empty List"
	}
	for ; currentNode.next != l.head; currentNode = currentNode.next { //loop until current node equals first node(circular)
		if currentNode.elem == elem {
			return "Found"
		}
	}
	return "Not Found"
}
