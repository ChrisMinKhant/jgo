package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, _ := fsnotify.NewWatcher()
	watcher.Add("/media/kaungminkhant/external-disk/jgo/.jgo")
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fetchFileData()
				fmt.Printf("Fetched event ::: %v\n", event)
				done <- true

			case errors := <-watcher.Errors:
				fmt.Printf("Fetched error ::: %v\n", errors)
				return
			}
		}
	}()

	<-done
}

func fetchFileData() {
	readFile, _ := os.ReadFile("/media/kaungminkhant/external-disk/jgo/.jgo")

	extractedValue := strings.Split(string(readFile), "/<")

	decodedString, _ := base64.StdEncoding.DecodeString(extractedValue[0])

	fmt.Printf("Fetched read file ::: %v\n", string(decodedString))
}
