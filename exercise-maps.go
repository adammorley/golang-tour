package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
    p := strings.Fields(s)
    for _, v := range p {
        m[v] += 1
    }
    return m
}

func main() {
    wc.Test(WordCount)
}
