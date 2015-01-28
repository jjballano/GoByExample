package main 

import "fmt"

func main() {
	//Three different types of for
	i := 0

	for i < 3 {
		fmt.Println("For without autoincrement and previous var declaration", i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println("For with autoincrement and not previous var declaration", j)		
	}

	k := 0
	for ;k < 3; k++ {
		fmt.Println("For with autoincrement and previous var declaration", k)
	}

	for {
		fmt.Println("Infinite for (break ahead)")
		break
	}

}