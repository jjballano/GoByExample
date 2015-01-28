package main 

import "fmt"

func main() {
	
	m := make(map[string]int)
	m["key1"] = 9
	m["key2"] = 5

	fmt.Println("Map: ",m)

	fmt.Println("Element key1: ",m["key1"])

	fmt.Println("lenght:", len(m))

	delete(m, "key2")
    fmt.Println("key2 removed:", m)

    value, prs := m["key2"]
    fmt.Print("Value ", value)
    fmt.Println("       but was found? ", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map initialized:", n)
}