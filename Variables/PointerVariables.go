// File Name: PointerVariables.go

package main

import "fmt"

func Main() {
    x := 42
    p := &x
    fmt.Println(*p) // Output: 42
}
