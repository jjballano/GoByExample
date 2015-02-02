package main 

import "fmt"

func main() {
	fmt.Println("\"make\" creates a new channel ")
	messages := make(chan string)

	go func(){
		fmt.Println("Program waits until channel is written")
		messages <- "ping"
	}()

	msg := <- messages
	fmt.Println("Program waits until channel is read")
	fmt.Println(msg)
}