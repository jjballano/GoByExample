package main 

import ("fmt"
		"errors")

func f1(arg int) (int, error) {
	if arg < 0 {
		return -1, errors.New("Can't work with numbers lower than 0")
	}

	return arg - 1, nil
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("Value => %d     Text => %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg < 0 {
		return -1, &argError{arg, "can't work with it"}
	}

	return arg - 1, nil
}

func main() {
	fmt.Println("By convention, errors are the last return value and have type error")
	fmt.Println("It’s possible to use custom types as errors by implementing the Error() method on them")

	fmt.Println("------")
	fmt.Println("Function will return (result and nil) or (-1 and error.New(...))")
	for _,i := range[]int{-1, 0, 5}{
		if r, e := f1(i); e != nil{
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
			if r == -1 {
				fmt.Println("Even worked with a result == -1")
			}
		}
	}

	fmt.Println("------")
	fmt.Println("Function will return (result and nil) or (-1 and &argError{...})")
	for _,i := range[]int{-1, 0, 5}{
		if r, e := f2(i); e != nil{
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked", r)
			if r == -1 {
				fmt.Println("Even worked with a result == -1")
			}
		}
	}

	fmt.Println("------")
	_, e := f2(-1)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }

}	
