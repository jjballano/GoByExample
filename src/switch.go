package main

import "fmt"

func main() {

    i := 2
    fmt.Print("write ", i, " as ")
    
    switch i {
        case 1:
            fmt.Println("one")
        case 2, 3:
            fmt.Println("two or three")
        default:
            fmt.Println("Other values")
    }

    fmt.Println("Not break needed")

    fmt.Println("Not only number switch available")
}