package main 

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func check(e error){
	if (e != nil){
		panic(e)
	}
}

func main() {
	fmt.Println("Read files.")
	fmt.Println("/tmp/dat file must exists")
	fmt.Println("ioutil.ReadFile(file) read the file to memory")
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("/tmp/dat")
	defer f.Close()
    check(err)
    fmt.Println("os.Open(file) returns a pointer to the file",f)

    fmt.Println("Read function")
    b1 := make([]byte,3) //It is going to read 3 bytes
    n1,err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))

    fmt.Println("Seek + Read functions")
    o2, err := f.Seek(5, 0)
    check(err)
    b2 := make([]byte, 4)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

    fmt.Println("ReadAtLeast function")
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    fmt.Println("'Rewind' function, Seek(0,0)")
    _, err = f.Seek(0, 0)
    check(err)

    fmt.Println("bufio.NewReader")
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))
}