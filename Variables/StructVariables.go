// File Name: StructVariables.go

package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func Main() {
    p := Person{Name: "Md", Age: 25}
    fmt.Println(p) // Output: {Md 25}
}
