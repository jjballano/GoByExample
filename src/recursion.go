package main 

import "fmt"

func main() {
	fmt.Println("Fact(4) = ",fact(4))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}