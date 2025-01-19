package main

import (
    "fmt"
)

func main() {
    var input int
    fmt.Printf("input number: ")
    fmt.Scan(&input)
    if input > 0 {
        fmt.Println("positive")
    } else if input < 0 {
        fmt.Println("negative")
    } else {
        fmt.Println("zero")
    }
}

