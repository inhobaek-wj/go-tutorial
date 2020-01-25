package main

import "fmt"

func makeSquares(slice []int) {
    fmt.Printf("address in makeSquares: %v\n", &slice[0])
    fmt.Printf("address in makeSquares: %v\n", &slice[1])

    for i, elem := range slice {
        slice[i] = elem * elem
    }
}

// changing the slice itself is not working,
// but changing items in the slice is working.
func split(slice []int) {
    fmt.Printf("address in split: %v\n", &slice[0])
    fmt.Printf("address in split: %v\n", &slice[1])

    slice = slice[1:5]
    fmt.Printf("%v : \n", slice)

    for i, elem := range slice {
        slice[i] = elem * 2
    }
}

func main() {
    s := []int{0,1,2,3,4,5,6,7,8,9}
    fmt.Printf("address in main: %v\n", &s[0])
    fmt.Printf("address in main: %v\n", &s[1])
    fmt.Printf("%v: \n", s)

    split(s)
    fmt.Printf("%v: \n", s)

    makeSquares(s)
    fmt.Printf("%v: \n", s)

    a := []int{1,2,3}
    b := a
    b[1] = 10
    fmt.Printf("%v\n", a)
    fmt.Printf("%v\n", &a[0])
    fmt.Printf("%v\n", &b[0])
    c := a[:]
    fmt.Printf("%v\n", &c[0])


    d := make([]int,2)
    d[0] = 100
    fmt.Printf("%v\n", d[0])
    fmt.Printf("%v\n", &d[0])
    fmt.Printf("length: %v\n", len(d))
    fmt.Printf("capacity: %v\n", cap(d))

    // e := []int{}
    // fmt.Printf("%v\n", &e[0]) // panic: runtime error: index out of range [0] with length 0.
    // fmt.Printf("length: %v\n", len(e))
    // fmt.Printf("capacity: %v\n", cap(e))

    d = append(d,1,2,3,4,5)
    d[0] = 200
    fmt.Printf("%v\n", d[0])
    fmt.Printf("%v\n", &d[0])
    fmt.Printf("length: %v\n", len(d))
    fmt.Printf("capacity: %v\n", cap(d))

    d = append(d,1,2,3,4,5)
    fmt.Printf("%v\n", d[0])
    fmt.Printf("%v\n", &d[0]) // addresses on same index are different.
    fmt.Printf("length: %v\n", len(d))
    fmt.Printf("capacity: %v\n", cap(d))


    f := make([]int, 2)
    fmt.Printf("%v\n", &f[0])
    fmt.Printf("length: %v\n", len(f))
    fmt.Printf("capacity: %v\n", cap(f))

    // using unpack operator.
    f = append(f, []int{1,2,3,4,5}...) // same with not using unpack operator.
    fmt.Printf("%v\n", &f[0])
    fmt.Printf("length: %v\n", len(f))
    fmt.Printf("capacity: %v\n", cap(f))

    fmt.Println("----------------------------")
    o := append(f,1,2,3) // since o's capacity exceed f's capacity, new array will be created.
    fmt.Printf("%v\n", &f[0])
    fmt.Printf("length: %v\n", len(f))
    fmt.Printf("capacity: %v\n", cap(f))

    fmt.Printf("o: %v\n", &o[0])
    fmt.Printf("o's length: %v\n", len(o))
    fmt.Printf("o's capacity: %v\n", cap(o))

    p :=
        append(f[:], 1,2,3) // since o's capacity exceed f's capacity, new array will be created.
    fmt.Printf("%v\n", &f[0])
    fmt.Printf("length: %v\n", len(f))
    fmt.Printf("capacity: %v\n", cap(f))

    fmt.Printf("p: %v\n", &p[0])
    fmt.Printf("p's length: %v\n", len(p))
    fmt.Printf("p's capacity: %v\n", cap(p))

    q := append(f[:2], 1,2,3) // q's capacity doesn't exceed f's capacity, it references f.
    fmt.Printf("%v\n", &f[0])
    fmt.Printf("length: %v\n", len(f))
    fmt.Printf("capacity: %v\n", cap(f))

    fmt.Printf("q: %v\n", &q[0])
    fmt.Printf("q's length: %v\n", len(q))
    fmt.Printf("q's capacity: %v\n", cap(q))

    // sine Go does not provide fancy functions for slice,
    // we need to hack.

    // shift.
    g := []int{1,2,3,4,5}
    h := g[1:]
    fmt.Printf("%v\n", h)

    // pop.
    g = []int{1,2,3,4,5}
    h = g[:len(g) - 1]
    fmt.Printf("%v\n", h)

    // delete i = 3.
    deleteIdx := 3
    l := make([]int, len(g))
    copy(l,g)
    l = append(l[:deleteIdx-1], l[deleteIdx:]...)
    fmt.Printf("g: %v\n", g)
    fmt.Printf("l: %v\n", l)
    // fmt.Printf("m: %v\n", m)

    fmt.Printf("l[0]: %v\n", &l[0])
    // fmt.Printf("m[0]: %v\n", &m[0])

    // copy slice.
    i := []int{1,2,3,4,5}
    j := make([]int,5)
    copy(j,i)
    fmt.Printf("%v\n", &i[0])
    fmt.Printf("%v\n", &j[0])
    fmt.Printf("%v\n", i)
    fmt.Printf("%v\n", j)

    j[0] = 9
    fmt.Printf("%v\n", i)
    fmt.Printf("%v\n", j)

    k := make([]int, len(i))
    copy(k,i)
    fmt.Printf("%v\n", k)

    r := []int{1,2,3,4,5}
    fmt.Printf("%v\n", &r[0])
    ss := r[:]
    fmt.Printf("s: %v\n", &ss[0])
}
