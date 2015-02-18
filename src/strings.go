package main 

import ("fmt"
		s "strings"
		"os")

var p = fmt.Println

func main() {

	p("- Functions -")
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test","te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace (all):   ", s.Replace("foo", "o", "0", -1))
    p("Replace (first):   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
    //not part of strings but related with them
    p("Len: ", len("hello"))
    p("Char:", "hello"[1])
	
}