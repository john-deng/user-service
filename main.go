package main

import "fmt"

func foo()  {
    fmt.Println("Hello, foo")
}

func main() {
    fmt.Println("Hello, world")
    fmt.Println("Hello, go")
    fmt.Print("Hello again")
    foo()
}
