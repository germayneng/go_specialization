package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	// create empty int slice of length 3
	var sli = make([]int, 3)
	var exit bool
	//while
	for exit == false {
		var userInput string
		fmt.Println("Please input an integer:")
		_, _ = fmt.Scan(&userInput)

		if userInput == "X" {
			exit = true
		} else {
			convert, _ := strconv.Atoi(userInput)
			sli = append(sli, convert)

		}
		clip := sli[3:] // clip from index 3 onwards
		sort.Ints(clip) // sort the slice
		fmt.Println(clip)

	}
}
