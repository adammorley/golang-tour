package main

import "fmt"

func foo() int {
    z := 4
    return z
}

func main() {
    var i, j  int = 1, 2
    k := 3
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
    fmt.Println(foo())
}
