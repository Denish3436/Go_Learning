package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(10, 15)
    fmt.Printf("The sum is: %d\n", sum)
}

package main

import (
    "fmt"
    "errors"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero 🚫")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("The result is: %.2f\n", result)
    }

    // Try with a zero divisor
    _, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    }
}
