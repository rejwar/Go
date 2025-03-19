
package main

import "fmt"

func main() {
    var data interface{}
    data = "Hello"
    fmt.Println("String:", data)
    data = 123
    fmt.Println("Integer:", data)
}
