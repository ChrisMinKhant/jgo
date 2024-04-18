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
		classBuilder := extractFunctionInfo(&readFile, entry.Name())
		fileName := entry.Name()

		createFile(&fileName, classBuilder)
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

			fmt.Printf("Fetched splitted function info ::: %v \n", splittedFunctionInfo)

			goLibBuilder.SetMethodParameter(extractFunctionParameter(&splittedFunctionInfo))
		}
	}

	return goLibBuilder
}

func extractFunctionParameter(splittedFunctionInfo *[]string) *string {
	parameters := ""
	tempSplittedFunctionInfo := *splittedFunctionInfo

	for index, functionInfo := range tempSplittedFunctionInfo {

		returnedDataType := dataTypeMapper(&functionInfo)

		if returnedDataType != "" {
			if !strings.Contains(tempSplittedFunctionInfo[index-1], "(") {
				parameters += returnedDataType + " " + tempSplittedFunctionInfo[index-1] + ","
				continue
			}

			parameters += returnedDataType + " " + strings.Split(tempSplittedFunctionInfo[index-1], "(")[1] + ","

		}

	}

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

func dataTypeMapper(functionInfo *string) string {
	returnedDataType := ""

	for dataType := range DataTypeMap {
		if strings.Contains(*functionInfo, dataType) {
			returnedDataType = DataTypeMap[dataType]
			return returnedDataType
		}
	}

	return ""
}

func createFile(fileName *string, classBuilder *goLibBuilder) {

	createdFile, error := os.Create("/media/kaungminkhant/New Volume/jgo/src/main/java/com/jgo/framework/gorunner/golib/" + *fileName + ".java")

	if error != nil {
		panic(error.Error())
	}

	createdFile.WriteString(*classBuilder.Build())

	defer createdFile.Close()
}
