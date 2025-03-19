package main

import (
    "fmt"
    "os"
)

func main() {
    os.Setenv("APP_ENV", "development")
    fmt.Println("APP_ENV:", os.Getenv("APP_ENV"))
}
