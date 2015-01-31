package main 

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int{
	return r.width * r.height
}

func (r rect) perim() int{
	return r.width * 2 + r.height * 2
}

func (r *rect) randomMethodWithExtraParam(anotherArea int) int{
	return r.area() - anotherArea
}

func main() {
	myRect := rect{width:10, height:20}

	fmt.Println("Area ('pointer' method): ", myRect.area())
	fmt.Println("Perim (no 'pointer' method): ", myRect.perim())
	fmt.Println("Method that receives an extra param: ", myRect.randomMethodWithExtraParam(20))
}