package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	readDir, _ := os.ReadDir("../")

	for _, entry := range readDir[:len(readDir)-1] {
		readFile, _ := os.ReadFile("../" + entry.Name() + "/" + entry.Name() + ".go")
		extractFunctionInfo(&readFile, entry.Name())
	}
}

func extractFunctionInfo(inputFile *[]byte, fileName string) *goLibBuilder {

	// TO DO
	// set class name
	// set method name
	// set method parameter
	// set method return type

	goLibBuilder := NewGoLibBuilder()

	fetchedValueFromFile := string(*inputFile)

	splittedFileValue := strings.Split(fetchedValueFromFile, "\n")

	for index, singleString := range splittedFileValue {

		if strings.Contains(singleString, "//export") {

			splittedString := strings.Split(singleString, " ")

			goLibBuilder.SetClassName(&fileName)
			goLibBuilder.SetMethodName(&splittedString[1])

			splittedFunctionInfo := strings.Split(splittedFileValue[index+1], " ")

			goLibBuilder.SetMethodParameter(extractFunctionParameter(&splittedFunctionInfo))
		}
	}

	return goLibBuilder
}

func extractFunctionParameter(splittedFunctionInfo *[]string) *string {
	parameters := ""
	tempSplittedFunctionInfo := *splittedFunctionInfo

	for index, functionInfo := range tempSplittedFunctionInfo {

		dataTypeMapper(&functionInfo)
		if strings.Contains(functionInfo, "int") {
			if !strings.Contains(tempSplittedFunctionInfo[index-1], "(") {
				parameters += "int " + tempSplittedFunctionInfo[index-1] + ","
				continue
			}

			parameters += "int " + strings.Split(tempSplittedFunctionInfo[index-1], "(")[1] + ","
		}

	}

	// fmt.Printf("Fetched method signature ::: %v\n", len(functionSignature))

	if parameters != "" {
		parameters = parameters[:len(parameters)-1]
	}

	return &parameters
}

var DataTypeMap = map[string]string{
	"int":     "int",
	"int8":    "int",
	"int16":   "int",
	"int32":   "int",
	"int64":   "int",
	"bool":    "bool",
	"string":  "String",
	"byte":    "byte",
	"float32": "float",
	"float64": "float",
}

func dataTypeMapper(functionInfo *string) {
	dataTypeChannel := make(chan string)

	go func() {
		for dataType := range DataTypeMap {
			fmt.Printf("Fetched data type ::: %v\n", dataType)
		}
	}()

	<-dataTypeChannel
}

// createdFile, error := os.Create("/home/kaungminkhant/jgo-framework/src/main/java/com/jgo/framework/gorunner/golib/" + *fileName + ".java")

// 	if error != nil {
// 		panic(error.Error())
// 	}

// 	createdFile.WriteString(*goLibBuilder.Build())

// 	defer createdFile.Close()
