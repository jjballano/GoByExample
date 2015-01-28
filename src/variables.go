package main 

import "fmt"

func main() {
	
	var aString string = "initial"
	fmt.Println("var aString string = \"initial\" => ",aString)
	fmt.Println("--")
	var anotherString = "I don't need to specify the type when a variable is declared"
	fmt.Println("var anotherString = xxx => ", anotherString)
	fmt.Println("--")
	fmt.Println("ANOTHERstring => Error. Variables are case sensitive")
	fmt.Println("var anyString STRING => Error. Types are case sensitive")
	fmt.Println("--")
	fmt.Println("var aString \n ab=\"hi\" => Compilation error.")
	fmt.Println("var aString string \n ab=\"hi\" => Works.")
	fmt.Println("--")
	var aInt = 3
	fmt.Println("aInt = ",aInt)
	fmt.Println("//I don't need to specify the type when a variable is declared")
	fmt.Println("--")
	var a,b int = 1,2
	fmt.Println("var a,b int = 1,2 => You can declare multiple variables at once. a = ",a)
	fmt.Println("--")
	fmt.Println("I must use all variables, otherwise I'll get a compilation error so b=",b)
	fmt.Println("--")
	var zeroValue int
	var emptyString string
	var initBool bool
	fmt.Println("Variables declared without a corresponding initialization are zero-valued => Integer: zeroValue=",zeroValue)
	fmt.Println("Variables declared without a corresponding initialization are zero-valued => String: emptyString=",emptyString)
	fmt.Println("Variables declared without a corresponding initialization are zero-valued => Boolean: initBool=",initBool)
	fmt.Println("--")
	f := "f := \"Value\" => Better way to declare variables"
	fmt.Println(f)

}