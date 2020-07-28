/*
This assignment gives the example to
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// define struct
	type Names struct {
		fname string
		lname string
	}

	// slice containing Names
	var sli = make([]Names, 0)

	// read file.txt
	var fileName string
	fmt.Println("Enter file name containing names...")
	_, _ = fmt.Scan(&fileName) //give filename's address so Scan can store it
	data, error := os.Open(fileName)

	if error != nil {
		log.Fatal(error)

	}

	// using bufio
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		// split the strings via space
		strArray := strings.Split(scanner.Text(), " ")

		// append the names to the struc
		var temp Names
		temp.fname = strArray[0]
		temp.lname = strArray[1]
		// store in slice
		sli = append(sli, temp)

	}

	// iterate through and print using for each
	for _, names := range sli {
		fmt.Println(names.fname)
		fmt.Println(names.lname)

	}
}
