package main

import "fmt"

func main() {
    var p *int
    fmt.Println(p)

    i := 10
    p = &i
    fmt.Println(p)
    fmt.Println(i)
    fmt.Println(*p)
}