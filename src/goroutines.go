package main 

import "fmt"

func f(from string) {
	for i:=0; i < 3; i++ {
		fmt.Println(from,":",i)
	}
	
}

func main() {
	
	fmt.Println("A goroutine is a lightweight thread of execution")

	fmt.Println("'go f(\"in a goroutine\")' will invoke the function in a go routine. It will execute concurrently with the calling one")
	go f("Direct")

	fmt.Println("You can also start a goroutine for an anonymous function call")

	go func(msg string) {
		fmt.Println(msg)
	}("goroutine for an anonymous funcion call")

	//Our two function calls are running asynchronously in separate goroutines now, so execution falls through to here. This Scanln code requires we press a key before the program exits
	var input string
    fmt.Scanln(&input)
    fmt.Println("done")

}