package main 

import ("fmt"
		"regexp"
		"bytes")

func main() {
	
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println("- FindString")
	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))
	fmt.Println("- FindAllString")
	fmt.Println(r.FindAllString("peach punch notthis pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2)) //number of matches limited 
	fmt.Println("- FindStringSubmatch")
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println("- FindAllStringSubmatch")
	fmt.Println(r.FindAllStringSubmatchIndex("peach notthis pinch", -1))

	fmt.Println("- Others")
	fmt.Println(r.Match([]byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch") //When creating constants with regular expressions you can use the MustCompile variation of Compile. A plain Compile won’t work for constants because it has 2 return values
    fmt.Println(r)
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))

}