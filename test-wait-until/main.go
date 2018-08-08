package main

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"fmt"
	"time"
)

func main() {
	shutdown := make(chan struct {})
	go wait.Until(print, time.Second*2, shutdown)

	select {
	case <-time.After(time.Second * 10):
		close(shutdown)
	case <-time.After(time.Second*2):
		print()
	case <-shutdown:

	}
}

func print() {
	fmt.Println(time.Now().Unix())
}