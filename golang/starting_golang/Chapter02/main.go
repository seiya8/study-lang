package main

import (
    "fmt"

    "Chapter02/animals"
)

func main() {
    fmt.Println(AppName())

    fmt.Println(animals.ElephantFeed())
    fmt.Println(animals.MonkeyFeed())
    fmt.Println(animals.RabbitFeed())
}
