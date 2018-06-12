package main

import (
    "io"
    "os"
    "strings"
)

/*
    takes an input character as an integear (ascii)
    and enciphers it with rot13.  later, upgrade to unicode
*/
func rot13(in int) int {
    if (in < 65 || in > 90 && in < 97 || in > 122) {
        return in
    }
    if (in <= 97 && in + 13 > 90) { // wrap
        return in + 13 - 90 + 64
    } else if (in + 13 > 122) { // wrap
        return in + 13 - 122 + 96
    } else {
        return in + 13
    }
}

type rot13Reader struct {
    r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
    n, err := r.r.Read(b)
    if err != nil {
        return n, err
    }
    for i := 0; i < n; i++ {
        b[i] = byte(rot13(int(b[i])))
    }
    return n, nil
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
