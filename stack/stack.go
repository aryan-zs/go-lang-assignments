package stack

import (
    "fmt"
)

type Stk struct {
    data []int
}

func (s *Stk) Push(a int) {
    s.data = append(s.data, a)
}


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
