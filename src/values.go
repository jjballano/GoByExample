package main 

import "fmt"

func main() {
	
	fmt.Println("go" + "lang")

	fmt.Println("3-5=", 3-5)
	fmt.Println("7/5=", 7/5)
	fmt.Println("7.0/5.0=", 7.0/5.0)

	fmt.Println("true && false ", true && false)
	fmt.Println("true || false ", true || false)
	fmt.Println("!true", !true)

	fmt.Println("3 + 'Text' => Error (mismatched types int and String)")

	fmt.Println("'Single quote' is _NOT_ allowed ")
	fmt.Println("a variable can have nil value")
}