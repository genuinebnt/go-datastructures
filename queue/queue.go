package queue

import "fmt"

//Queue contains items that are added at end and removed from beginning
type Queue struct {
	items []int
}

//Push add element to end of list
func (q *Queue) Push(elem int) {
	q.items = append(q.items, elem)
}

//Show shows all elements in queue
func (q *Queue) Show() []int {
	return q.items
}

//Pop removes first element of queue
func (q *Queue) Pop() []int {
	if len(q.items) == 0 {
		fmt.Println("Queue empty.")
		return q.items
	}
	items := q.items[1:]
	q.items = items
	return items
}

//Len returns length of queue
func (q *Queue) Len() int {
	return len(q.items)
}
