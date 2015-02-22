package main 

import ("flag"
		"fmt")

func main() {
	wordPtr := flag.String("word", "foo", "a string") //name, default value, short description
	var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var") //same as previuos way

	//flag.String return a string pointer, not a string value
	numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    flag.Parse() //Once all flags are declared, call flag.Parse() to execute the command-line parsing.

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())

}