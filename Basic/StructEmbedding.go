package main

import "fmt"

type Animal struct {
    Name string
}

type Dog struct {
    Animal
    Breed string
}

func main() {
    d := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Golden Retriever"}
    fmt.Println("Dog's name:", d.Name, "| Breed:", d.Breed)
}
