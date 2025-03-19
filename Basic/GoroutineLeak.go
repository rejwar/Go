package main

import "fmt"

func doWork(ch chan bool) {
    for range ch {
        fmt.Println("Doing work")
    }
}

func main() {
    ch := make(chan bool)
    go doWork(ch)
    close(ch)
}
