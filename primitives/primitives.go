package main

import (
	"fmt"
	"strconv"
)
func main() {

    var a bool = true
    fmt.Printf("%v, %T\n", a, a) // true, bool

    b := 1 == 2
    fmt.Printf("%v, %T\n", b, b) // false, bool

    var c bool
    fmt.Printf("%v, %T\n", c, c) // false, bool

    d := 42
    fmt.Printf("%v, %T\n", d, d) // 42, int

    // int8        -128 ~ 127
    // int16       -32 768 ~ 32 767
    // int32       -2 147 483 648 ~ 2 147 483 647
    // int64       - 9 223 372 036 854 755 808 ~ 9 223 372 036 854 775 807

    // uint8       0 ~ 255
    // uint16      0 ~ 65 535
    // uint32      0 ~ 4 294 967 295

    // var e uint8 = 500 // constant 500 overflows uint8
    // fmt.Printf("%v, %T\n", e, e)

    var f int = 10
    var g int8 = 3
    // fmt.Println(f + g) this is not working
    fmt.Println(f + int(g))

    h := 10 // 1010
    i := 3 // 0011
    fmt.Println(h & i) // 0010 = 2
    fmt.Println(h | i) // 1011 = 11
    fmt.Println(h ^ i) // 1001 = 9
    fmt.Println(h &^ i) // 0100 = 8

    j := 8 // 2^3
    fmt.Println(j << 3) // 2^3 * 2^3 = 2^6 = 64
    fmt.Println(j >> 3) // 2^3 / 2^3 = 2^0 = 1

    // float32       +-1.18E-38 ~ +-3.4E38
    // float64       +-2.23E-308 ~ +-1.8E308

    // when using initialize syntax on decimal value,
    // it will be initialize to float64.
    k := 3.14
    fmt.Printf("%v, %T\n", k, k) // 3.14, float64

    var l complex64 = 1 + 2i
    fmt.Printf("%v, %T\n", l, l) // (1+2i), complex64
    fmt.Printf("%v, %T\n", real(l), real(l)) // 1, float32
    fmt.Printf("%v, %T\n", imag(l), imag(l)) // 2, float32

    var m complex128 = complex(5, 12)
    fmt.Printf("%v, %T\n", m, m) // (5+12i), complex128

    // text, string type
    // string is a slice of bytes.
    // in go, string stands for any utf-8 character.
    n := "this is a string"
    fmt.Printf("%v, %T\n", n[2], n[2]) // 105, uint8
    fmt.Printf("%v, %T\n", string(n[2]), n[2]) // i, uint8

    // n[2] = "u" cannot assign to n[2], becuase string is read-only slice.

    p := []byte(n)
    fmt.Printf("%v, %T\n", p, p) // [116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8

    // in go, rune represent for any utf-32 character.
    // rune is type alias of int32.
    // use single qoutes, and it is character literal.
    r := 'a'
    fmt.Printf("%v, %T\n", r, r) // 97, int32

    o := "이건 한글"
    fmt.Printf("%v, %T\n", o, o) // 이건 한글, string
    fmt.Printf("%v, %T\n", o[1], o[1]) // 157, uint8
    fmt.Printf("%v, %T\n", string(o[1]), string(o[1])) // , string -> character is broken.

    s := []rune(o)
    fmt.Printf("%v, %T\n", s, s) // [51060 44148 32 54620 44544], []int32
    fmt.Printf("%v, %T\n", s[1], s[1]) // 44148, int32

    u := []byte(o)
    fmt.Printf("%v, %T\n", u, u) // [236 157 180 234 177 180 32 237 149 156 234 184 128], []uint8

    v := "건"
    fmt.Printf("%v, %T\n", v, v) // 건, string

    w := []byte(v)
    fmt.Printf("%v, %T\n", w, w) // [234 177 180], []uint8

    t := strconv.QuoteRune(rune(o[1]))
    fmt.Printf("%v, %T\n", t, t) // '\u009d', string
    fmt.Printf("%v, %T\n", rune(o[1]), rune(o[1])) // '\u009d', string

    t = strconv.QuoteRune(s[1])
    fmt.Printf("%v, %T\n", t, t) // '건', string


}
