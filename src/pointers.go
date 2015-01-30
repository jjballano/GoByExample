package main 

import "fmt"

func main() {
	value := 3
	fmt.Println("Init value ",value)
	fmt.Println("&value ",&value)

	fmt.Println("----")

	valueParameterFunction(value)
	fmt.Println("valueParameterFunction(value)",value)

	pointerParameterFunction(&value)
	fmt.Println("pointerParameterFunction(&value)",value)
}

func valueParameterFunction(param int){
	param = 0
}


func pointerParameterFunction(pointer *int){
	*pointer = 0
}