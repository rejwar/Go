package main

import "fmt"

func main() {
    day := "Wednesday"
    switch day {
    case "Monday", "Tuesday", "Wednesday":
        fmt.Println("Weekday")
    case "Saturday", "Sunday":
        fmt.Println("Weekend")
    default:
        fmt.Println("Invalid day")
    }
}
