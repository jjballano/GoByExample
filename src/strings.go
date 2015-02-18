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

    point := point{1,2}
    fmt.Printf("%v\n", point)
    fmt.Printf("%#v\n", point) //prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value.
	fmt.Printf("%T\n", point)
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%d\n", "123")
	fmt.Printf("%b\n", 14) //binary
	fmt.Printf("%c\n", 33) //character corresponding to the given integer
	fmt.Printf("%x\n", 456) //hex encoding
	fmt.Printf("%x\n", "hex this")
	fmt.Printf("%f\n", 78.9) //floats. basic decimal formatting 
	//%e and %E format the float in (slightly different versions of) scientific notation
	fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)
    fmt.Printf("%s\n", "\"string\"") //basic string printing
    //To double-quote strings as in Go source, use %q.
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%q\n", "string")
	fmt.Printf("%p\n", &point) // print a representation of a pointer
	fmt.Printf("|%10d|%4d|\n", 12, 345) //control the width and precision of the resulting figure.By default the result will be right-justified and padded with spaces
	fmt.Printf("|%-10d|%-4d|\n", 12, 345) //the result will be left-justified and padded with spaces
	fmt.Printf("|%10.4f|%6.1f|\n", 1.2, 3.45)
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
    p(s)

    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}

type point struct {
	x, y int
}