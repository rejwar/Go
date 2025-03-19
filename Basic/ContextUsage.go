package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Operation timed out")
    case <-ctx.Done():
        fmt.Println("Context canceled:", ctx.Err())
    }
}
