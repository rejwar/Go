package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    jsonString := `{"name":"Md.","age":25}`
    var p Person
    json.Unmarshal([]byte(jsonString), &p)
    fmt.Println(p)
}
