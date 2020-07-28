package main

import "fmt"

// by using slice, we do not need to do the pointer thing for reference. By default it is referencing

// demo helper function that alters the first index
func helper(x []int) {
	x[0] = x[0] * 2
}

func main() {

	// make a slice (instead of array)
	var x = make([]int, 0)
	x = append(x, 100)

	//before
	fmt.Println(x)

	//after
	helper(x)
	fmt.Println(x)

}
