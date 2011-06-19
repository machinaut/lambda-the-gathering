// ajr.go, first crack at ICFP 2011 Programming Competition

package main

import (
    "fmt"
)

var (
    u   = make([]Slot, 256) // proponent
    υ   = make([]Slot, 256) // opponent
)

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
