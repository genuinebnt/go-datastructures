package stack

import "fmt"

//Stack struct that holds a list of elements and behaves like a stack
type Stack struct {
	items []int
}

//Len returns length of stack
func (s *Stack) Len() int { return len(s.items) }

//Pop removes last element of list and returns it
func (s *Stack) Pop() int {
	topIdx := len(s.items) - 1

	if len(s.items) == 0 {
		fmt.Println("stack underflow")
		return -1
	}

	item := s.items[topIdx]
	s.items = s.items[0:topIdx]
	return item
}

//Push adds element to end of stack struct
func (s *Stack) Push(elem int) []int {
	s.items = append(s.items, elem)
	return s.items

}
