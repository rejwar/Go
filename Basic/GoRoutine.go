package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 3; i++ {
        fmt.Println(s)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    go say("Hello")
    say("World")
}
