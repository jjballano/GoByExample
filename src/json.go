package main 

import ("encoding/json"
		"fmt"
		"os")

var println = fmt.Println

type Response1 struct{
	Page int
	Fruits []string
}

type Response2 struct{
	Page int `json:"the page"`
	Fruits []string `json:"the fruits"`
}

func main() {
	bolB, _	:= json.Marshal(true)
	println(string(bolB))
	intB, _ := json.Marshal(1)
	println(string(intB))
	fltB, _ := json.Marshal(2.34)
    println(string(fltB))
    strB, _ := json.Marshal("gopher")
    println(string(strB))

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    println(string(mapB))

    res1D := &Response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    println(string(res1B))

    res2D := &Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    println(string(res2B))

    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    var dat map[string]interface{}

    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    println(dat)

    num := dat["num"].(float64) //casting the value to float64
    fmt.Println(num)

    println (dat["strs"])
    strs := dat["strs"].([]interface{})
    println(strs[0])
    str1 := strs[0].(string)
    println(str1)

    str := `{"the page": 1, "the fruits": ["apple", "peach"]}`
    res := &Response2{}
    json.Unmarshal([]byte(str), &res)
    println(res)
    println(res.Fruits[0])

    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}