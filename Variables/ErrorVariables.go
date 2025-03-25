// File Name: ErrorVariables.go

package main

import (
    "errors"
    "fmt"
)

func Main() {
    err := errors.New("Something went wrong!")
    fmt.Println(err) // Output: Something went wrong!
}
