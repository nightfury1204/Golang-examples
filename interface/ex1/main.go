package ex1

import "fmt"

type AInterface interface {
	Hello()
}

type BInterface interface {
	AInterface
	Bye()
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

type B struct {
	// 'a AInterface' will be failed to implement BInterface
	AInterface
	goodBye string
}

func NewB(a AInterface) *B {
	return &B{
		AInterface:a,
		goodBye: "Goodbye",
	}
}

func (b *B) Bye() {
	fmt.Println(b.goodBye)
}

func main()  {
	var a AInterface
	a = NewA()

	var b BInterface
	b = NewB(a)

	b.Hello()
	b.Bye()
}

