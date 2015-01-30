package main 

import "fmt"

func main() {
	nextInt := intSeq()

	fmt.Println("nextInt() => ", nextInt())
	fmt.Println("nextInt() => ", nextInt())
	fmt.Println("nextInt() => ", nextInt())

	fmt.Println("Restarting count...")
	newNextInt := intSeq()

	fmt.Println("newNextInt() => ", newNextInt())

}

func intSeq() func() int {
	i:=0
	return func() int {
		i += 1
		return i
	}
}