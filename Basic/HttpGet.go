package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    response, err := http.Get("https://example.com")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer response.Body.Close()
    body, _ := ioutil.ReadAll(response.Body)
    fmt.Println("Response:", string(body))
}
