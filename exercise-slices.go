package main

import "fmt"
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    r := make([][]uint8, dy)
    for i := 0; i < dy; i++ {
        r[i] = make([]uint8, dx)
        for j := 0; j < dx; j++ {
            r[i][j] = uint8((i+j)/2)
        }
    }
    return r
}

func main() {
    r := Pic(5, 5)
    fmt.Printf("%T\n", r)
    pic.Show(Pic)
}
