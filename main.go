package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

func watch(watcher *fsnotify.Watcher, file *string, verbose *bool) {
	err := watcher.Add(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ignore %s\n", *file)
		return
	}
	if *verbose == true {
		fmt.Fprintf(os.Stderr, "watch %s\n", *file)
	}
}

func main() {
	verbose := flag.Bool("verbose", false, "more output (default: false)")
	flag.Parse()

	watcher, watcher_err := fsnotify.NewWatcher()
	if watcher_err != nil {
		panic(watcher_err)
	}
	defer watcher.Close()

	for _, arg := range flag.Args() {
		if arg != "-" {
			watch(watcher, &arg, verbose)
			continue
		}

		stdin := bufio.NewScanner(os.Stdin)
		for stdin.Scan() {
			line := stdin.Text()
			watch(watcher, &line, verbose)
		}
	}

	fmt.Fprintf(os.Stderr, "[%4d files] wait for changes...\n", len(watcher.WatchList()))
	event, ok := <-watcher.Events

	if ok == false {
		panic(fmt.Errorf("event not ok"))
	}

	fmt.Fprintf(os.Stderr, "[%s] %s\n", event.Op, event.Name)

	os.Exit(0)
}
