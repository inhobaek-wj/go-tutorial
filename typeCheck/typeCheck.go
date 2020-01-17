package main

import (
    "fmt"
)

type myStruct struct{}

func main() {
    i := myStruct{}
    i.test()

    fmt.Printf("%T\n", i)

    isTest(i)
}

type myType interface {
    test()
}

func (m myStruct) test() {
    fmt.Println("abc")
}

func isTest(t interface{}) {
    switch t.(type) {
    case int:
        fmt.Println("Im int")
    case myType:
        fmt.Println("Im myType")
    case myStruct:
        fmt.Println("Im myStruct")
    default:
        fmt.Println("default")
    }
}
