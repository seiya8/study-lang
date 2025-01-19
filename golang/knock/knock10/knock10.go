package main

import (
    "fmt"
)

func main() {
    var input int
    fmt.Printf("input number: ")
    fmt.Scan(&input)
    if input < 0 {
        input *= -1
    }
    fmt.Println("absolute value is ", input)
}

