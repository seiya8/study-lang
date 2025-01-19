package main

import (
    "fmt"
)

func main() {
    var n, m int
    fmt.Printf("input 1st number: ")
    fmt.Scan(&n)
    fmt.Printf("input 2nd number: ")
    fmt.Scan(&m)
    fmt.Println("和: ", n + m)
    fmt.Println("差: ", n - m)
    fmt.Println("積: ", n * m)
    fmt.Println("商: ", n / m, ", 余り: ", n % m)
}

