package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x int = 42
    t := reflect.TypeOf(x)
    fmt.Println("Type:", t)
    v := reflect.ValueOf(x)
    fmt.Println("Value:", v)
}
