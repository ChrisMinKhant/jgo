package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io"
	"log"
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
	readFileMessage, _ := readFile("/mnt/edisk/jgo/.secret")

	secretAndIv := strings.Split(*readFileMessage, "\n")

	log.Printf("Fetched secret and iv ::: %v\n", secretAndIv[0])

	aes, error := aes.NewCipher([]byte(secretAndIv[0]))

	if error != nil {
		log.Printf("Fetched found error at cipher ::: %v\n", error)
	}
	log.Printf("Fetched AES ::: %v\n", aes)

	gcm, _ := cipher.NewGCM(aes)

	log.Printf("Fetched nonce size ::: %v\n", gcm.NonceSize())

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
