// File Name: DataStructures.go

package main

import "fmt"

func Main() {
    Numbers := []int{1, 2, 3, 4}
    Numbers = append(Numbers, 5)
    fmt.Println(Numbers) // Output: [1, 2, 3, 4, 5]
}
