package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// sorts a series of integers
func sortInts(sli []int, c chan []int) {
	fmt.Println("Sorting... ", sli)
	sort.Ints(sli)
	c <- sli
}

// converts a string array to int array
func convertStringArray(sa []string) []int {
	var intAr = []int{}
	for _, i := range sa {
		j, err := strconv.Atoi(strings.TrimSpace(i))
		if err != nil {
			panic(err)
		}
		intAr = append(intAr, j)
	}
	return intAr
}

// find the min between two numbers
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	// start a new reader to capture user input
	in := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter a comma seperated series of ints: ")
	usernums, _ := in.ReadString('\n')
	stringArr := strings.Split(usernums, ",")
	// call convertStringArray to convert each element to an int
	intArr := convertStringArray(stringArr)
	// define a 2d array to store seperate chunks
	chunks := make([][]int, 0)
	// chunk array into seperate batches
	d := math.Ceil(float64(len(intArr)) / 4.0)
	for i := 0; i < len(intArr); i += int(d) {
		batch := intArr[i:min(i+int(d), len(intArr))]
		chunks = append(chunks, batch)
	}
	// define a channel with 4 object buffer
	c := make(chan []int, 4)
	// start 4 goroutines for sorting
	go sortInts(chunks[0], c)
	go sortInts(chunks[1], c)
	go sortInts(chunks[2], c)
	go sortInts(chunks[3], c)
	// retrieve 4 values from the buffer
	var sortedBatch1 = <-c
	var sortedBatch2 = <-c
	var sortedBatch3 = <-c
	var sortedBatch4 = <-c
	// print the seperate sorted batches
	fmt.Println("Sorted Batch #1 ", sortedBatch1)
	fmt.Println("Sorted Batch #2 ", sortedBatch2)
	fmt.Println("Sorted Batch #3 ", sortedBatch3)
	fmt.Println("Sorted Batch #4 ", sortedBatch4)

	// merge sorted slices
	x := []int{}
	x = append(sortedBatch1, sortedBatch2...)
	x = append(x, sortedBatch3...)
	x = append(x, sortedBatch4...)

	// sort final slice
	go sortInts(x, c)
	var finalBatch = <-c

	// print the final slice
	fmt.Println("Final Sorted Array #1 ", finalBatch)

}
