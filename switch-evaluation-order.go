package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Whe's Saturday?")
    today := time.Now().Weekday()
    switch time.Saturday {
        case today + 0:
            fmt.Println("today")
        case today + 1:
            fmt.Println("Tomorrow")
        default:
            fmt.Println("Long time!")
    }
}
