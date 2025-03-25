// File Name: FunctionVariables.go

package main

import "fmt"

func Add(a, b int) int {
    return a + b
}

func Main() {
    result := Add(10, 20)
    fmt.Println(result) // Output: 30
}
