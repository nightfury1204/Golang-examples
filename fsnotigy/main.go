package main

import (
	"github.com/appscode/kutil/tools/fsnotify"
	"path/filepath"
	"fmt"
	"os"
	"os/signal"
	"log"

	fs "github.com/fsnotify/fsnotify"
)
func main() {
	ExampleNewWatcher()

	w := fsnotify.Watcher{
		WatchDir: filepath.Dir("/home/ac/go/src/Golang-examples/fsnotigy/data/a.txt"),
		Reload: func() error {
			fmt.Println("file changed")
			return nil
		},
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	stopCh := make(chan struct{})

	err := w.Run(stopCh)
	fmt.Println(err)

	<-signalChan
	close(stopCh)
}

func ExampleNewWatcher() {
	watcher, err := fs.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fs.Write == fs.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/home/ac/go/src/Golang-examples/fsnotigy/data/a.txt")
	if err != nil {
		log.Fatal(err)
	}
	<-done
	fmt.Println("done")
}