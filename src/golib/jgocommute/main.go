package main

var JGO_FILE = "/mnt/edisk/jgo/.jgo"

func main() {
	jgoWatcher := NewJgoWatcher()
	jgoWatcher.Watch()
}
