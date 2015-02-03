package main 

import ("fmt"
		"time")

func main() {

	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C //blocks on the timer channel C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	//You can cancel the timer before it expires
	go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    } 
}