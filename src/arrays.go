package main 

import "fmt"

func main() {
	
	var b [5]string
	fmt.Println("Array with exaclty 5 string values and not initialized: ", b)

	var a [5]int
	fmt.Println("Array with exaclty 5 int values and not initialized: ", a)

	a[4] = 100
    fmt.Println("Same array of int values after change position 4:", a)
    fmt.Println("get position 4:", a[4])

    fmt.Println("len(a):", len(a))

    c := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Array with 5 int values initialized:", c)

    d := [2][3]int{{1, 2}, {4, 5, 6}}
    fmt.Println("Array multidimension [2][3] int values initialized:", d)
}