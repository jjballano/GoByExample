package main 

import "fmt"

func main() {
	fmt.Println("singleReturnValue(1, 2)=",singleReturnValue(1, 2))

	fmt.Println("-----")
	
	label, sum := multipleReturnValues(1, 2)
	fmt.Println("multipleReturnValues(1, 2), first result => ",label)
	fmt.Println("multipleReturnValues(1, 2), second result => ",sum)

	fmt.Println("-----")

	fmt.Println("We could use _,sum:=multipleReturnValues(1, 2) if we don't care about the first result")
}

func singleReturnValue(num1 int, num2 int) int{
	return num1 + num2
}

func multipleReturnValues(num1 int, num2 int)(string, int){
	sum := num1 + num2
	return "The sum of the params is",sum
}