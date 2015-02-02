package main 

import "fmt"

func main() {
	jobs := make (chan int, 5)
	done := make (chan bool)

	go func() {
		for {
			//the 'more' will be false if jobs has been closed and all values in the channel have already been received
			j, more := <-jobs
			if more {
				fmt.Println("received jobs: ",j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 10; j++ {
		jobs <- j
		fmt.Println("sent job",j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<- done
}