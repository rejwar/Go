package main

import "fmt"

func main() {
    defer fmt.Println("Goodbye!")
    fmt.Println("Hello!")
}
