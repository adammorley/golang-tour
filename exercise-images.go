package main

import "image"
import "image/color"
import "golang.org/x/tour/pic"

type AnImage struct{
    w int
    h int
    r uint8
    g uint8
    b uint8
    a uint8}

func (a AnImage) ColorModel() color.Model {
    return color.RGBAModel
}

func (a AnImage) Bounds() image.Rectangle {
    return image.Rect(0, 0, a.w, a.h)
}

func (a AnImage) At(x, y int) color.Color {
    return color.RGBA{a.r, a.g, a.b, a.a}
}

func main() {
    m := AnImage{w: 25, h: 25, r: 10, g: 5, b: 6, a: 9}
    pic.ShowImage(m)
}
