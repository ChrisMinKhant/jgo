package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

type JgoDispatcher struct {
	readFileMessage *string
	readFileByte    *[]byte
}

func NewJgoDispatcher(readFileMessage *string, readFileByte *[]byte) *JgoDispatcher {
	return &JgoDispatcher{
		readFileMessage: readFileMessage,
		readFileByte:    readFileByte,
	}
}

func (jgoDistpatcher *JgoDispatcher) dispatch() {

	pushMessage := decodeMessage(jgoDistpatcher.readFileMessage)

	fmt.Printf("Fetched push message ::: %v\n", pushMessage)

	jgoDistpatcher.commit()
}

func (jgoDistpatcher *JgoDispatcher) commit() {
	openedFile, _ := os.OpenFile(JGO_FILE, os.O_RDWR, 0644)

	openedFile.Seek(int64(len(*jgoDistpatcher.readFileByte)), io.SeekCurrent)

	openedFile.Truncate(0)

	openedFile.Close()
}

func decodeMessage(encodedMessage *string) *[]string {
	fmt.Printf("Fetched encoded message ::: %v\n", *encodedMessage)
	messageSlice := strings.Split(*encodedMessage, "|")
	decodedMessage, _ := base64.StdEncoding.DecodeString(messageSlice[len(messageSlice)-2])

	var pushMessage []string

	messageSlice = strings.Split(string(decodedMessage), "|")

	for _, singleMessage := range messageSlice {
		decodedMessage, _ = base64.StdEncoding.DecodeString(singleMessage)
		pushMessage = append(pushMessage, string(decodedMessage))
	}

	fmt.Printf("Fetched push message ::: %v\n", pushMessage)

	return &pushMessage
}
