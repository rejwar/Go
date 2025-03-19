package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    file.WriteString("Hello, Go!")
    fmt.Println("File written successfully!")
}
