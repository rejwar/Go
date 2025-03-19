package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    t := reflect.TypeOf(User{})
    for i := 0; i < t.NumField(); i++ {
        fmt.Println("Tag for", t.Field(i).Name, "is", t
