// File Name: FunctionsAndMethods.go

package main

import "fmt"

type Person struct {
    Name string
}

func (p Person) Greet() string {
    return "Hello, " + p.Name
}

func Main() {
    P := Person{Name: "Md"}
    fmt.Println(P.Greet())
}
