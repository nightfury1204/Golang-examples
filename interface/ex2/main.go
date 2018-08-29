package main

import "fmt"

type AInterface interface {
	Hello()
}

type A struct {
	greetings string
}

func NewA() *A {
	return &A{
		greetings:"Good morning",
	}
}

func (a *A) Hello() {
	fmt.Println(a.greetings)
}

func main(){
	var a AInterface
	a = NewA()
	b := a.(*A)
	fmt.Println(b.greetings)
}

