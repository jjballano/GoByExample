package main 

import ("fmt"
		"sort")

func main() {
	strs := []string{"C","a","B", "z"}
	sort.Strings(strs)
	fmt.Println("Sort is case sensitive in strings")
	fmt.Println("It changes the array, so sorting is in-place")
	fmt.Println("List sorted (string): ",strs)

	ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted? ", s)

    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println("Sorted by length ",fruits)
}		

type ByLength []string //In order to sort by a custom function in Go, we need a corresponding type. Here we’ve created a ByLength type that is just an alias for the builtin []string type

//We implement sort.Interface - Len, Less, and Swap
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}


