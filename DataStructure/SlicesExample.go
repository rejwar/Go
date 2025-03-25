// File Name: SlicesExample.go

package main

import "fmt"

func Main() {
    Numbers := []int{1, 2, 3}           // Dynamic size
    Numbers = append(Numbers, 4, 5)    // Adding elements
    fmt.Println(Numbers)               // Output: [1 2 3 4 5]
}
