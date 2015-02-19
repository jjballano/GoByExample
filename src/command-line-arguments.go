package main

import ("os"
		"fmt")

func main() {

	//First value is the path to the program
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:] //holds the arguments to the program.

    arg := os.Args[3] //Get individual arg
    
    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}