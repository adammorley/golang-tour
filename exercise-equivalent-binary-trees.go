package main

import "fmt"
import "golang.org/x/tour/tree"

func WalkTree(t *tree.Tree, c chan int) {
    Walk(t, c)
    close(c)
}

func Walk(t *tree.Tree, c chan int) {
    if t.Left != nil {
        Walk(t.Left, c)
    }
    c <- t.Value
    if t.Right != nil {
        Walk(t.Right, c)
    }
}


func printer(c chan int) {
    for i := range c {
        fmt.Println(i)
    }
}

func Same(t1, t2 *tree.Tree) bool {
    c1 := make(chan int)
    c2 := make(chan int)
    go WalkTree(t1, c1)
    go WalkTree(t2, c2)
    var j int
    var ok bool
    for i := range c1 {
        j, ok = <-c2
        if ok == false {
            panic("should not get here")
        }
        if i != j {
            return false
        }
    }
    return true
}

func main() {
    var t1, t2 *tree.Tree
    t1 = tree.New(1)
    t2 = tree.New(1)
    fmt.Println(Same(t1, t2))
}
