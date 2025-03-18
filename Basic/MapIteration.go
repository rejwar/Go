package main

import "fmt"

func main() {
    capitals := map[string]string{"Bangladesh": "Dhaka", "USA": "Washington D.C."}
    for country, capital := range capitals {
        fmt.Println(country, "->", capital)
    }
}
