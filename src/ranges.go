package main 

import "fmt"

func main() {
	
	nums := []int{1,2,3,4}
	sum := 0

	for _,num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	fmt.Println("-----")

    fmt.Println("for with index.")
	for i, num := range nums {
        fmt.Print(" Value of num:", num)
        fmt.Println(" Value of index:", i)
    }
    fmt.Println("-----")
    kvs := map[string]string{"a":"apple", "b":"banana"}
    fmt.Println("For for map. ")
    for key, value := range kvs {
    	fmt.Printf("Key=>%s  Value=>%s\n", key, value)
    }

    fmt.Println("-----")
    fmt.Println("range on strings iterates over Unicode code points. Example with \"golan\" ")
    for index, char := range "golan" {
        fmt.Print("Char=>", char)
        fmt.Println("  Index=>",index)
    }
}