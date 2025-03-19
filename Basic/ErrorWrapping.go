package main

import (
    "errors"
    "fmt"
)

func main() {
    originalErr := errors.New("original error")
    wrappedErr := fmt.Errorf("an additional context: %w", originalErr)
    fmt.Println("Error:", wrappedErr)
}
