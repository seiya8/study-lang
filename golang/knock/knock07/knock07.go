package main

import (
    "fmt"
)

func main() {
    var input int
    fmt.Printf("input number: ")
    fmt.Scan(&input)
    if input == 0 {
        fmt.Println("zero")
    } else {
        fmt.Println("not zero")
    }
}

