package main 

import ("fmt"
		"math/rand")

func main() {
	p := fmt.Print
	pln := fmt.Println
//It is not really random but pseudorandom. Always returns the same value. It is deterministic
	p(rand.Intn(100), ",")
	p(rand.Intn(100))
	pln()

	pln(rand.Float64())

	p((rand.Float64()*5)+5, ",")
    p((rand.Float64() * 5) + 5)
    pln()

    s1 := rand.NewSource(42)
    r1 := rand.New(s1)

    p(r1.Intn(100),",")
    p(r1.Intn(100))
    pln()

    s2 := rand.NewSource(42) //Using the same seed, it will generate the same results
    r2 := rand.New(s2)
    p(r2.Intn(100), ",")
    p(r2.Intn(100))
    pln()
}		