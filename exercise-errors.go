package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt struct {
    Msg string
    Val float64
}

func (e *ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", e.Val)
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return -1, &ErrNegativeSqrt{"nope", x}
    }
    z := float64(1)
    for i := 1; i < 5; i++ {
        z -= (z*z - x) / (2*z)
    }
    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
    x, e := Sqrt(-2)
    fmt.Println(x)
    fmt.Println(e)
}
