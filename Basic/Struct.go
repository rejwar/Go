package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    person := Person{Name: "Md.", Age: 25}
    fmt.Println(person)
}
