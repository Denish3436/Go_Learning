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
// Tags for Struct Fields
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

package main

import "fmt"

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {

    p := person{name: name}
    p.age = 42
    return &p
}


//example 
func main() {

    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30})

    fmt.Println(person{name: "Fred"})

    fmt.Println(&person{name: "Ann", age: 40})

    fmt.Println(newPerson("Jon"))

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)

    dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
}

