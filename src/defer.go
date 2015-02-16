package main 

import (
	"os"
	"fmt"
)

//"defer" is like "finally" in Java
func main() {
	
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(filename string) *os.File {
	fmt.Println("Creating file ")
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(file *os.File) {
	fmt.Println("Writing in the file")
	fmt.Fprintln(file, "data")
}

func closeFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}