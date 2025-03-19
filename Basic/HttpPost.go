package main

import (
    "bytes"
    "fmt"
    "net/http"
)

func main() {
    data := []byte(`{"key":"value"}`)
    response, err := http.Post("https://example.com", "application/json", bytes.NewBuffer(data))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer response.Body.Close()
    fmt.Println("Response Status:", response.Status)
}
