package main

import "fmt"

func main() {
    capitals := map[string]string{
        "Bangladesh": "Dhaka",
        "India":      "New Delhi",
    }
    fmt.Println("Capital of Bangladesh is:", capitals["Bangladesh"])
}
