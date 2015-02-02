package main 

import "fmt"

func main() {
	fmt.Println("\"make\" creates a new channel ")
	messages := make(chan string)

	//Channels buffered
	buffered := make(chan string, 2)

	go func(){
		fmt.Println("Program waits until channel is written, because it expects to read something")
		messages <- "ping"

		buffered <- "First"
		buffered <- "Second"
	}()

	msg := <- messages
	msg += <- buffered
	msg += <- buffered

	fmt.Println(msg)
}