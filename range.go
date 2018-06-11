package main

import (
    "fmt"
    "math"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }

    for i := 0; i < 7; i++ {
        fmt.Printf("2**%d = %d\n", i, int(math.Pow(float64(2), float64(i))))
    }
}
