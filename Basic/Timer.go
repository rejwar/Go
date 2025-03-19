package main

import (
    "fmt"
    "time"
)

func main() {
    timer := time.NewTimer(3 * time.Second)
    <-timer.C
    fmt.Println("Timer expired")
}
