package main

import "C"
import "fmt"

func main() {

}

type TransferredJsonData struct {
	Data string
}

//export TestSecond
func TestSecond(transferredJsonData string) {
	fmt.Printf("Fetched transfereed data ::: %v\n", transferredJsonData)
}
