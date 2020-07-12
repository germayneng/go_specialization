package main

import (
	"fmt"
	"strconv"
)

//m3 1.1 arrays -----
// we initialize the fixed size array of len 5
func demoArray() {
	var xArray [5]int
	fmt.Println(xArray) // [0 0 0 0 0]
	fmt.Printf("Capacity of array: %d \n", len(xArray))
}

func demoConvert() {

	fmt.Println(strconv.Atoi("X"))
}

// main func ---------
func main() {

	//demoArray()
	demoConvert()

}
