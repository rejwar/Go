package main

import "fmt"

func main() {
    add := func(a, b int) int {
        return a + b
    }
    fmt.Println("Sum:", add(3, 7))
}
