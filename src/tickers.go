package main 

import ("fmt"
		"time")

func main() {
	//Tickers are for when you want to do something repeatedly

	ticker := time.NewTicker(time.Millisecond * 500)
	go func(){
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}()

	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}