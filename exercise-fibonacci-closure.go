package main

import "fmt"

func fibDP() func() int {
    var f [100]int
    f[0] = 0
    f[1] = 1
    var n int
    n = -1
    return func() int {
        n++
        if n > 1 {
            f[n] = f[n-1] + f[n-2]
        }
        return f[n]
    }
}

func fibMC() func() int {
    var n1, n2, n, v int
    n2 = 0
    n1 = 1
    n = -1
    return func() int {
        n++
        if n == 0 {
            return n2
        } else if n == 1 {
            return n1
        } else {
            v = n2 + n1
            n2 = n1
            n1 = v
            return v
        }
    }
}


func main() {
    f := fibDP()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
    f = fibMC()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
