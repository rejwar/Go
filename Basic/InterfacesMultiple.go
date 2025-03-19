package main

import "fmt"

type Shape interface {
    Area() float64
}

type Square struct {
    Side float64
}

func (s Square) Area() float64 {
    return s.Side * s.Side
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    shapes := []Shape{Square{Side: 4}, Circle{Radius: 5}}
    for _, shape := range shapes {
        fmt.Println("Area:", shape.Area())
    }
}
