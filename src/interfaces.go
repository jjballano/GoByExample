package main 

import ("fmt"
		"math")

type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}

func (s square) perim() float64 {
	return s.width * 2 + s.height * 2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return math.Pi * c.radius * 2
}

func measure(g geometry) {
	fmt.Println("Element: ",g)
	fmt.Println("Area: ",g.area())
	fmt.Println("Perim: ",g.perim())
}

func main() {
	
	circle := circle{radius:10}
	square := square{width:10, height:20}

	measure(circle)
	measure(square)
	fmt.Println("You don't declare in strut definition that it implements any interface, you just implements the methods of the interface")
}