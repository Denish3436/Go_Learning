//Functions
package main

import "fmt"
func greet() {
    fmt.Println("Hello, 👋 Welcome to Go functions.")
}

func main() {
    greet() // Call the function
}

//Functions with Parameters
package main

import "fmt"

func greetUser(name string) {
    fmt.Printf("Hello, %s! 🎉 Welcome to Go functions.\n", name)
}

func main() {
    greetUser("Denish")  // Try changing the name
    greetUser("Alex")
}
