package main

import "fmt"

func sum(x int) int {
    if x == 1{
        return 1
    } else {
        return x + sum(x-1)
    }
}

func test_sum() int {
    if sum(5) != 15 {
        panic("sum does not work")
    }
    return 0
}

func main() {
    fmt.Println(sum(5))
    test_sum()
    fmt.Println(sum(5))
}
