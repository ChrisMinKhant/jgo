package main

import "fmt"
import "C"

func main() {
}

//export Test
func Test(firstValue int, secondValue int) {
	thirdValue := firstValue + secondValue
	fmt.Printf("Hello this is from gotest ::: %v \n", thirdValue)
}
