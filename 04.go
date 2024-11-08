package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(10, 15)
    fmt.Printf("The sum is: %d\n", sum)
}

