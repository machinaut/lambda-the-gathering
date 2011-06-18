package main

type Card struct {
    Id    int    //
    Short byte   // Shortname
    Name  string // long name
    Desc  string // long name
}

var Deck = []Card{
    Card{0, 'z', "zero", "λx.0"},
    Card{1, 'I', "I", "λx.x"},
    Card{2, 's', "succ", "λx.x+1"},
    Card{3, 'd', "dbl", "λx.x+x"},
    Card{4, 'g', "get", "λi.ρ[i]"},
    Card{5, 'p', "put", "λx.I"},
    Card{6, 'S', "S", "λf.λg.λx.fx(gx)"},
    Card{7, 'K', "K", "λx.λy.x"},
    Card{8, 'i', "inc", "λi.I v[i]←v[i]+1"},
    Card{9, 'd', "dec", "λi.I ν[255-i]←ν[255-i]-1"},
    Card{10,'a', "attack", "λi.λj.λn.fx(gx)"},
}
