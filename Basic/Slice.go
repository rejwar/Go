package main

import "fmt"

func main() {
    colors := []string{"red", "green", "blue"}
    colors = append(colors, "yellow")
    fmt.Println("Colors:", colors)
}
