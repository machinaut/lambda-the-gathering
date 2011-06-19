// ajr.go, first crack at ICFP 2011 Programming Competition

package main

import (
    "fmt"
)

type Slot struct {
    f   uint8
    val uint16
    vit int32
}

var (
    u   = make([]Slot, 256) // proponent
    υ   = make([]Slot, 256) // opponent
)

func (s Slot) alive() bool {
    return s.vit > 0
}

func (s Slot) fun(v ...int) int {
    switch s.f {
    case 0:
        return int(s.val)
    case 1:
        return v[0]
    }
    return 0
}

func ValidSlotNumber(n int) bool {
    return n > 0 && n < 256
}

func LeftApplication() {
}
func RightApplication() {
}

func main() {
    fmt.Println("AJR")
    for _, c := range Deck {
        fmt.Println(c)
    }
}
