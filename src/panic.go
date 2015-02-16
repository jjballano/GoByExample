package main 

import ("os"
		"fmt")

func main() {
	panic ("a problem")

	fmt.Println("code not executed")
	
	_,error := os.Create("/tmp/file")
	if error != nil {
		//It will throw an exception if there is a problem creating a file

		panic (error)
	}
}