package main

import (
    "fmt"
    "image"
)

func main() {
    n := image.Rect(0, 0, 100, 100)
    describe(n)
    m := image.NewRGBA(n)
    fmt.Println(n.Bounds())
    fmt.Println(m.Bounds())
    fmt.Println(m.At(0,0).RGBA())
}
