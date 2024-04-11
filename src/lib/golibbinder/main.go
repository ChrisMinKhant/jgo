package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	goLibBuilder := NewGoLibBuilder()

	readDir, _ := os.ReadDir("../")

	for _, entry := range readDir {
		readFile, _ := os.ReadFile("../" + entry.Name() + "/" + entry.Name() + ".go")

		fetchedValue := string(readFile)

		stringArray := strings.Split(fetchedValue, " ")

		for index, singleString := range stringArray {
			fmt.Printf("Fetched parts of program ::: %v \n", singleString)

			if strings.Contains(singleString, "//export") {
				fmt.Printf("Fetched single string ::: %v\n", stringArray[index+1])
			}
		}

		createdFile, error := os.Create("/home/kaungminkhant/jgo-framework/src/main/java/com/jgo/framework/gorunner/golib/" + entry.Name() + ".java")

		if error != nil {
			panic(error.Error())
		}

		createdFile.WriteString(*goLibBuilder.BuildGoLibJavaClass())

		defer createdFile.Close()
	}
}
