package main

import (
	"encoding/json"
	"fmt"
)

type resp struct {
	A int `json:"a,omitempty"`
	B string `json:"b,omitempty"`

}

func main() {
	d1 := `
{
"a":1223
}
`
	d2 := `
{
"b":"hi"
}
`
	a1 := resp{B:"yoo"}
	b1 := resp{A:7}

	err := json.Unmarshal([]byte(d1), &a1)
	if err!=nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println(a1)
	}

	err = json.Unmarshal([]byte(d2), &b1)
	if err!=nil {
		fmt.Println("err: ", err)
	}else {
		fmt.Println(b1)
	}

	fmt.Println("--------------------------------------")

	d3 := `
{
"a":1223,
"bbb":"hi"
}
`
	a1 = resp{B:"yoo"}

	err = json.Unmarshal([]byte(d3), &a1)
	if err!=nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println(a1)
	}

	fmt.Println("--------------------------------------")

	d4 := `
{
"a":1223,
"bb":"hi"
}
`
	a1 = resp{}

	err = json.Unmarshal([]byte(d4), &a1)
	if err!=nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println(a1)
	}
}
