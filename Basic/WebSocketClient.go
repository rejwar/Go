package main

import (
    "fmt"
    "golang.org/x/net/websocket"
)

func main() {
    ws, err := websocket.Dial("ws://example.com/socket", "", "http://example.com/")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Fprintln(ws, "Hello WebSocket!")
    var msg string
    fmt.Fscan(ws, &msg)
    fmt.Println("Received:", msg)
}
