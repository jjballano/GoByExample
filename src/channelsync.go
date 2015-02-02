package main 

import ("fmt"
		"time")

func worker(done chan bool){
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	fmt.Println("Channels are used to synchronize execution across goroutines")
	done := make(chan bool, 1)

	go worker(done)

	<- done
}