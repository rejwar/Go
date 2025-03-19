package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    content, err := ioutil.ReadFile("output.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("File content:", string(content))
}
