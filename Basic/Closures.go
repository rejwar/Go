package main

import "fmt"

func main() {
    counter := func() func() int {
        x := 0
        return func() int {
            x++
            return x
        }
    }()

    fmt.Println(counter())
    fmt.Println(counter())
    fmt.Println(counter())
}
