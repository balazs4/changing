package main

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, watcher_err := fsnotify.NewWatcher()
	if watcher_err != nil {
		panic(watcher_err)
	}
	defer watcher.Close()

	for _, arg := range os.Args[1:] {
		err := watcher.Add(arg)
		if err != nil {
			panic(err)
		}
	}

	fmt.Fprintf(os.Stderr, "wait for changes...\n")
	event, ok := <-watcher.Events

	if ok == false {
		panic(fmt.Errorf("event not ok"))
	}

	fmt.Fprintf(os.Stderr, "[%s] %s\n", event.Op, event.Name)

	os.Exit(0)
}
