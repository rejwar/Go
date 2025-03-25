// File Name: LinkedListExample.go

package main

import "fmt"

type Node struct {
    Value int
    Next  *Node
}

func Main() {
    Head := &Node{Value: 10}
    Head.Next = &Node{Value: 20}
    Head.Next.Next = &Node{Value: 30}

    Current := Head
    for Current != nil {
        fmt.Println(Current.Value) // Output: 10, 20, 30
        Current = Current.Next
    }
}
