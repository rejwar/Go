package main

import "fmt"

func main() {
    ch1, ch2 := make(chan string), make(chan string)

    go func() { ch1 <- "Hello" }()
    go func() { ch2 <- "World" }()

    select {
    case msg1 := <-ch1:
        fmt.Println("Received from channel 1:", msg1)
    case msg2 := <-ch2:
        fmt.Println("Received from channel 2:", msg2)
    }
}
