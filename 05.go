package main

import "fmt"

// Define a struct called 'Person'
type Person struct {
    Name string
    Age  int
}

func main() {
    // Create a new Person
    var p Person
    p.Name = "Denish"
    p.Age = 21
    fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)
}

