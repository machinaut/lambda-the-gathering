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

func (c Card) String() string {
    return fmt.Sprintf("%c %02d %s %s", c.Short, c.Id, c.Name, c.Desc)
}

func (c Card) Z() {
}
func (c Card) I() {
}
func (c Card) s() {
}
func (c Card) D() {
}
func (c Card) g() {
}
func (c Card) p() {
}
func (c Card) S() {
}
func (c Card) K() {
}
func (c Card) i() {
}
func (c Card) d() {
}
func (c Card) a() {
}
func (c Card) h() {
}
func (c Card) c() {
}
func (c Card) r() {
}
func (c Card) z() {
}
