
package main

import "fmt"

func swap(a, b int) (int, int) {
    return b, a
}

func main() {
    x, y := swap(10, 20)
    fmt.Println("Swapped:", x, y)
}
