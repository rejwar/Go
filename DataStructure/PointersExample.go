// File Name: PointersExample.go

package main

import "fmt"

func Main() {
    X := 42
    P := &X
    *P = 50
    fmt.Println(X) // Output: 50
}
