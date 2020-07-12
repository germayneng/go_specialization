package main

import "fmt"

func main() {

	var UserInput float64
	fmt.Println("Please input a floating number...")
	_, _ = fmt.Scan(&UserInput)
	var truncated int32
	truncated = int32(UserInput)
	fmt.Printf("User input trancated value: %d \n", truncated)

}
