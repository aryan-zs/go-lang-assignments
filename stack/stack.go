package stack

import (
	"fmt"
)

// this is the stucture of the stack which inner data type is of sliced
type Stk struct {
	data []int
}

// this function is pushing the value  in the stack
func (s *Stk) Push(a int) {
	s.data = append(s.data, a)
}

// this function is poping the value from the stack and also removing that value from the stack
func (s *Stk) Pop() int {
	if len(s.data) == 0 {
		fmt.Println("Stack is empty")
		return -1 // default value
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

// this will only return the  top value of the stack but not remove it  from the stack
func (s *Stk) Top() int {
	if len(s.data) == 0 {
		fmt.Println("Stack is empty")
		return -1 // default value
	}
	val := s.data[len(s.data)-1]

	return val
}
