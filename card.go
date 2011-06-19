package main

import "fmt"

type Card struct {
    Id    int    // id number
    Short byte   // Shortname
    Name  string // long name
    Desc  string // long name
}

var Deck = []Card{
    Card{0, 'Z', "Zero", "λx.0"},
    Card{1, 'I', "I", "λx.x"},
    Card{2, 's', "succ", "λx.x+1"},
    Card{3, 'D', "dbl", "λx.x+x"},
    Card{4, 'g', "get", "λi.u[i]"},
    Card{5, 'p', "put", "λx.I"},
    Card{6, 'S', "S", "λf.λg.λx.fx(gx)"},
    Card{7, 'K', "K", "λx.λy.x"},
    Card{8, 'i', "inc", "λi.I v[i]←v[i]+1"},
    Card{9, 'd', "dec", "λi.I ν[255-i]←ν[255-i]-1"},
    Card{10, 'a', "attack", "λi.λj.λn.I v[i]←v[i]-n ν[255-j]←ν[255-j]-n*9/10"},
    Card{11, 'h', "help", "λi.λj.λn.I v[i]←v[i]-n v[j]←v[j]+n*11/10"},
    Card{12, 'c', "copy", "λi.u[i]"},
    Card{13, 'r', "revive", "λi.I v[i]←1"},
    Card{14, 'z', "zombie", "λi.λx.I υ[255-i]←x ν[255-i]←-1"},
}


func (c Card) f() {
    switch c.Short {
    case 'Z':
        c.Z()
    case 'I':
        c.I()
    case 's':
        c.s()
    case 'D':
        c.D()
    case 'g':
        c.g()
    case 'p':
        c.p()
    case 'S':
        c.S()
    case 'K':
        c.K()
    case 'i':
        c.i()
    case 'd':
        c.d()
    case 'a':
        c.a()
    case 'h':
        c.h()
    case 'c':
        c.c()
    case 'r':
        c.r()
    case 'z':
        c.z()
    }
}


func (c Card) Z() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
func (c Card) I() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
func (c Card) s() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
func (c Card) D() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
func (c Card) g() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
func (c Card) p() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
func (c Card) S() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
func (c Card) K() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
func (c Card) i() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
func (c Card) d() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
func (c Card) a() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
func (c Card) h() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
func (c Card) c() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
func (c Card) r() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
func (c Card) z() {
    fmt.Printf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}
