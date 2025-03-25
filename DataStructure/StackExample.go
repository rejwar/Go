// File Name: StackExample.go

package main

import "fmt"

type Stack struct {
    Elements []int
}

func (s *Stack) Push(Value int) {
    s.Elements = append(s.Elements, Value)
}

func (s *Stack) Pop() int {
    LastIndex := len(s.Elements) - 1
    Element := s.Elements[LastIndex]
    s.Elements = s.Elements[:LastIndex]
    return Element
}

func Main() {
    S := &Stack{}
    S.Push(10)
    S.Push(20)
    fmt.Println(S.Pop()) // Output: 20
}
