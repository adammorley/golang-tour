package main

import (
    "fmt"
    "strconv"
    "strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
    var sb strings.Builder
    for i, v := range ip {
        if (i > 0 && i < 4) {
            sb.WriteString(".")
        }
        s := strconv.Itoa(int(v))
        sb.WriteString(s)
    }
    return sb.String()
}

func main() {
    hosts := map[string]IPAddr{
        "loopback": {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    //for name, ip := range hosts {
    for _, ip := range hosts {
        //fmt.Printf("%v: %v\n", name, ip)
        fmt.Println(ip.String())
    }
}
