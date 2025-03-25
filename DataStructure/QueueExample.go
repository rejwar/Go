// File Name: QueueExample.go

package main

import "fmt"

type Queue struct {
    Elements []int
}

func (q *Queue) Enqueue(Value int) {
    q.Elements = append(q.Elements, Value)
}

func (q *Queue) Dequeue() int {
    Element := q.Elements[0]
    q.Elements = q.Elements[1:]
    return Element
}

func Main() {
    Q := &Queue{}
    Q.Enqueue(10)
    Q.Enqueue(20)
    fmt.Println(Q.Dequeue()) // Output: 10
}
