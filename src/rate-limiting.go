package main

import ("fmt"
		"time")

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<- limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("--------")

	burstyLimiter := make (chan time.Time, 5)
	for i:=0; i <5; i++ {
		burstyLimiter <- time.Now()
	}

	go func(){
		//This will constantly fill burstyLimiter every 200 millisecond
		for t := range time.Tick(time.Millisecond * 200){			
			fmt.Println("before execute: burstyLimiter <- t")
			burstyLimiter <- t
			fmt.Println("executed")
		}
	}()

	burstyRequests := make (chan int, 9)
	for i:=1; i <= 9; i++ {
		burstyRequests <- i
	} 
	close(burstyRequests)

	for req := range burstyRequests {
		<- burstyLimiter //this block execution until limiter has one "free element"
		fmt.Println("request", req, time.Now())
	}
	time.Sleep(3000 * time.Millisecond) //I want the goroutine to finish/be blocked
}