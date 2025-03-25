// File Name: ConcurrencyExample.go

package main

import (
    "fmt"
    "time"
)

func SayHello() {
    fmt.Println("Hello from Goroutine!")
}

func Main() {
    go SayHello()
    time.Sleep(1 * time.Second)
}
