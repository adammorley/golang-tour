package main

import "fmt"

type boolies struct {
    i int
    b bool
}

func main() {
    q := []int{2, 3, 5, 7, 11, 13}
    fmt.Println(q)

    r := []bool{true, false, true, true, false, true}
    fmt.Println(r)

    s := []boolies{
        {2, true},
        {3, true},
    }
    fmt.Println(s)
}
