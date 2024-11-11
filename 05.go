//Structs
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


package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

func main() {
    person := Person{Name: "Eve", Age: 32, Email: "eve@example.com"}

    jsonData, _ := json.Marshal(person)
    fmt.Println("JSON:", string(jsonData))
}
