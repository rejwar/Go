// File Name: MapsExample.go

package main

import "fmt"

func Main() {
    Capitals := map[string]string{
        "Bangladesh": "Dhaka",
        "India":      "New Delhi",
    }
    Capitals["Japan"] = "Tokyo" // Adding new key-value pair
    fmt.Println(Capitals)      // Output: map[Bangladesh:Dhaka India:New Delhi Japan:Tokyo]
}
