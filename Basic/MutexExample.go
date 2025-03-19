package main

import (
    "fmt"
    "sync"
)

var counter int
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
    mutex.Lock()
    counter++
    mutex.Unlock()
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go increment(&wg)
    }
    wg.Wait()
    fmt.Println("Final Counter Value:", counter)
}
