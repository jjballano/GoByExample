package main 

import ("fmt"
		"crypto/sha1")

func main() {
	s := "sha1 this string"
	hash := sha1.New()	

	hash.Write([]byte(s))

	bs := hash.Sum(nil)

	fmt.Println(s)
	fmt.Println(bs)
    fmt.Printf("In hex => %x\n", bs)	
}		