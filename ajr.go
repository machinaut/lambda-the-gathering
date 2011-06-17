// ajr.go, first crack at ICFP 2011 Programming Competition

package main

import (
    "fmt"
)

var (
    slot = make([]uint8, 100)
)

type slot struct {
    f       uint8
    val     uint16
    vit     int32
}

func (s slot) alive bool {
    return s.vit > 0;
}

func (s slot) f(v ...int) int {
    switch s.f {
        case 0:
            return s.val
        case 1:
            return v[0]
    }
}


func ValidSlotNumber(n int) bool {
    return n > 0 && n < 256
}


func main() {

}
