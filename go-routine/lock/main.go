package main

import (
	"sync"
	"fmt"
	"time"
	"strconv"
)

var (
	lock sync.Mutex
	greetings string
)

func greet(from, g string) {
	fmt.Println(from, "is here")
	lock.Lock()
	greetings = g
	fmt.Println(greetings,"from", from)
	select {
	case <-time.After(3*time.Second):
		fmt.Println("done from", from)
	}
	lock.Unlock()
}

func People(name string) {
	for {
		select {
		case <-time.After(30 * time.Second):
			return
		default:
			greet(name, "hi")
		}
	}
}

func main() {

	for i:=1;i<10;i++ {
		go People(strconv.Itoa(i))
	}

	select {
	case <-time.After(300 * time.Second):
	}
}