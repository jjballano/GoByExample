package main 

import ("fmt"
		"time"
		"sync/atomic"
		"runtime")

func main() {
	
	var ops uint64 = 0 //unsigned integer

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1) //autoincrement

				runtime.Gosched() //Allow other goroutines to proceed.
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops) //In order to safely use the counter while it’s still being updated by other goroutines, we extract a copy of the current value
	fmt.Println("ops:", opsFinal)


}