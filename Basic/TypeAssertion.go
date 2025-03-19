package main

import "fmt"

func main() {
    var i interface{} = "hello"
    s, ok := i.(string)
    if ok {
        fmt.Println("String value:", s)
    } else {
        fmt.Println("Not a string")
    }
}
