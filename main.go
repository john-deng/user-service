package main

import "fmt"

func foo()  {
    var count = 2017;
    fmt.Println("\nHello, foo", count)
}

func main() {
    fmt.Println("Hello, world")
    fmt.Println("Hello, go")
    fmt.Print("Hello again")
    foo()
}
