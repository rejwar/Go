package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(1 * time.Second)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()
    time.Sleep(5 * time.Second)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}
