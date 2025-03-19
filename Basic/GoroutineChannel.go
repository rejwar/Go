package main

import (
    "fmt"
)

func calculate(ch chan int, a, b int) {
    ch <- a + b
}

func main() {
    ch := make(chan int)
    go calculate(ch, 3, 4)
    result := <-ch
    fmt.Println("Result:", result)
}
