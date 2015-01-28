package main 

import "fmt"

const constant string = "Any constant value"

func main() {
	fmt.Println("Global constant value",constant)

	const anotherConst = "Another value"
	fmt.Println("Local constant value", anotherConst)
}