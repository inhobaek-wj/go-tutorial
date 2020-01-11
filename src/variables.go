/*
Visibility
- lower case first letter for package scope.
- upper case first letter to export.
- no private scope.

Naming conventions
- Pascal of camelCase
- Capitalize acronyms(HTTP, URL)

*/

package main

import (
        "fmt"
        "strconv"
)

// global scope variables.
var M int = 9

// package scope variables.
var l int = 42

// variable block.
var (
        name string = "inho"
        age float32 = 33.
)

// Local variable must be used.
// If not, complier will throw error.

// block scope variables.
func main() {

        // variable declaration.

        // var i int  = 42

        // var j int
        // j = 42

        k := 42

        // fmt.Println(i)
        // fmt.Println(j)
        // fmt.Println(k)

        fmt.Printf("%v, %T\n", k, k)


        fmt.Printf("%v, %T\n", l, l)
        var l int = 27 // variable shadowing.
        fmt.Printf("%v, %T\n", l, l)


        // explicit type conversion.
        var n int = 42
        fmt.Printf("%v, %T\n", n, n)

        var o float32
        o = float32(n)
        fmt.Printf("%v, %T\n", o, o)

        var p string
        p = string(n) // this is not working.
        fmt.Printf("%v, %T\n", p, p)

        var q string
        q = strconv.Itoa(n) // this is working.
        fmt.Printf("%v, %T\n", q, q)

}
