package main

import "fmt"

func main() {
    unbuffered := make(chan int)
    buffered := make(chan int, 2)

    go func() {
        unbuffered <- 1
    }()
    fmt.Println("Unbuffered:", <-unbuffered)

    buffered <- 1
    buffered <- 2
    fmt.Println("Buffered:", <-buffered, <-buffered)
}
