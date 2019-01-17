package main

import (
	"path/filepath"
	"fmt"
)

func main() {
	fmt.Println(filepath.Clean("hi/"))
	fmt.Println(filepath.Clean("hi"))
	fmt.Println(filepath.Clean("//hi//"))
	fmt.Println(filepath.Clean("/hi//"))
	fmt.Println(filepath.Clean("hello/hi//"))

	fmt.Println(filepath.Clean("hello/hi//")+"/")

	fmt.Println(filepath.Join("hi","/"))
	fmt.Println(filepath.Join("hi/","/"))

	fmt.Println(filepath.Join("hi","hello"))
	fmt.Println(filepath.Join("hi/","/hello"))
}
