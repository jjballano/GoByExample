package main 

import "fmt"

type person struct {
	name string	
	age int
	language string
}

func main() {

	fmt.Println(person{name:"Bob", age:21, language:"es"})

	fmt.Println(person{name:"Bob"})

	fmt.Println("You can create a variable with the same name as the struct => person := person{...}")
	person := person{name:"Paul", age:34, language:"en"}
	fmt.Println("Access to the property 'name' of the struct => person.name: ",person.name)

	anotherPerson := &person
	fmt.Println("You can create another variable with a pointer to another variable => anotherPerson := &person: ",anotherPerson.name)

	anotherPerson.age = 51
	fmt.Println("You can change a property value => anotherPerson.age = 51: ", anotherPerson.age)
	fmt.Println("Original variable value: ", person.age)

	noPointerVariable := person
	fmt.Println("noPointerVariable := person => ",noPointerVariable)
	noPointerVariable.age = 22
	fmt.Println("After changing noPointerVariable.age = 21, noPointerVariable = ", noPointerVariable)
	fmt.Println("After changing noPointerVariable.age = 21, person = ", person)

	fmt.Println("As excepted, only pointers change original value")

	fmt.Println("You can not create a new property in runtime, as expected")


}