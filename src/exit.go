package main

import "fmt"
import "os"

func main() {
	
	defer fmt.Println("!")
	fmt.Println("defer won't be run when using os.Exit")
	os.Exit(3)

}