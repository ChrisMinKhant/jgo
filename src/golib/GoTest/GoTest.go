package main

import "fmt"
import "C"

func main() {
}

//export Test
func Test(gibberish int, secondValue int, haha int) {
	fourthValue := gibberish + secondValue
	fmt.Printf("Hello this is from gotest ::: %v \n", fourthValue)
}
