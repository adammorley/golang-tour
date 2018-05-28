package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func splitSane(sum int) (int, int) {
    x := sum * 4 / 9
    y := sum - x
    return x, y
}

func main() {
    fmt.Println(split(17))
    fmt.Println(splitSane(17))
}
