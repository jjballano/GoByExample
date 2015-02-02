package main 

import "fmt"

func main() {
	fmt.Println("\"make\" creates a new channel ")
	messages := make(chan string)

	//Channels buffered
	buffered := make(chan string, 2)

	go func(){
		fmt.Println("Program waits until channel is written")
		messages <- "ping"

		buffered <- "First"
		buffered <- "Second"
	}()

	msg := <- messages
	msg += <- buffered
	msg += <- buffered

	fmt.Println("Program waits until channels are read")
	fmt.Println("For any reason, it doesnt need to wait for buffered channels")
	fmt.Println(msg)
}