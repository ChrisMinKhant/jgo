package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

type JgoWatcher struct {
	watcher    *fsnotify.Watcher
	dispatcher *JgoDispatcher
}

func NewJgoWatcher() *JgoWatcher {
	watcher, _ := fsnotify.NewWatcher()
	watcher.Add(JGO_FILE)

	return &JgoWatcher{
		watcher:    watcher,
		dispatcher: nil,
	}
}

func (jgoWatcher *JgoWatcher) Watch() {
	defer panicRecover(jgoWatcher)

	done := make(chan bool)

	go func() {
		for {
			select {
			case _, ok := <-jgoWatcher.watcher.Events:
				if !ok {
					log.Printf("Found non ok at jgo event watcher.")
					done <- ok
				}
				go jgoWatcher.respond()

			case Errors := <-jgoWatcher.watcher.Errors:
				fmt.Printf("Fetched errors ::: %v\n", Errors)
				done <- true
			}

		}
	}()

	panic(<-done)
}

func (jgoWatcher *JgoWatcher) respond() {
	readFileMessage, readFileByte := readFile(JGO_FILE)
	if *readFileMessage != "" {
		jgoWatcher.dispatcher = NewJgoDispatcher(readFileMessage, readFileByte)
		jgoWatcher.dispatcher.dispatch()
	}

}

func readFile(path string) (*string, *[]byte) {
	readValue, _ := os.ReadFile(path)
	fetchedMessage := string(readValue)
	return &fetchedMessage, &readValue
}

func panicRecover(jgoWatcher *JgoWatcher) {
	if r := recover(); r != nil {
		log.Printf("System has been reovered from panic ::: %v\n", r)
		jgoWatcher.Watch()
	}
}
