//loops
//for loop
package main

import "fmt"

func main() {
    countdown := 10
    for i := countdown; i > 0; i-- {
        fmt.Println(i)
    }
    fmt.Println("🎉 Happy New Year! 🎉")
}

//Infinite Loops - Handle with Care
package main

import (
    "fmt"
    "time"
)

func main() {
    i := 1
    for {
        fmt.Printf("Looping forever... %d\n", i)
        time.Sleep(1 * time.Second) // Pauses for a second
        i++
        
        if i > 5 { // Escape condition
            fmt.Println("Alright, enough looping. Breaking out! 🚪")
            break
        }
    }
}
