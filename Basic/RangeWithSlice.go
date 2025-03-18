package main

import "fmt"

func main() {
    nums := []int{1, 2, 3, 4}
    for i, v := range nums {
        fmt.Println("Index:", i, "Value:", v)
    }
}
