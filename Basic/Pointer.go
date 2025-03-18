package main

import "fmt"

func main() {
    a := 42
    b := &a
    fmt.Println("Value of a:", a)
    fmt.Println("Pointer to a:", b)
    fmt.Println("Value through pointer:", *b)
}
