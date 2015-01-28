package main 

import "fmt"

func main() {
	
	a := make([]int, 3)
	fmt.Println("make([]int, 3) => ",a)

	t := []string{"g", "h", "i"}
    fmt.Println("[]string{\"g\", \"h\", \"i\"} -> Slice initialized:", t)

	j := [][]string{{"a","b"}, {"h","j","k"}, {"h"}}
    fmt.Println("Slice multi-dimensional initialized:", j)

	fmt.Println("Slices are like array but support several more operations")

	s := make([]string, 2)
	s[0] = "hi"
	s[1] = "bye"
	s = append(s, "new", "another")
	fmt.Println("Slice after 'append' operation: ", s)

	c := make([]string, len(s))
    copy(c, s)
    fmt.Println("Copy of array 's':", c)

    fmt.Println("s[1:3] -> From 1 to 3 (not included) of 's':", s[1:3])
    fmt.Println("s[:3] -> From 0 to 3 (not included) of 's':", s[0:3])
    fmt.Println("s[2:] -> From 2 to the end of 's':", s[2:])
}