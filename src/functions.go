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

	fmt.Println("-----")
	fmt.Println("A Variadic Function can be called with any number of arguments ")
	fmt.Println("variadicFunction(1,2,3) => ",variadicFunction(1,2,3))
	fmt.Println("variadicFunction(1,2,3,4) => ",variadicFunction(1,2,3,4))
	nums := []int{4,5,6,7}
	fmt.Println("Can also be called with an array but adding ... after the param ")
	fmt.Println("variadicFunction(nums...) => ",variadicFunction(nums...))
}

func singleReturnValue(num1 int, num2 int) int{
	return num1 + num2
}

func multipleReturnValues(num1 int, num2 int)(string, int){
	sum := num1 + num2
	return "The sum of the params is",sum
}

func variadicFunction(nums ...int) int {
	sum := 0
	for _,num := range nums {
		sum += num
	}
	return sum
}