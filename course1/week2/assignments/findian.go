package main

import "fmt"

func main() {

	var UserInput string
	fmt.Println("Please input a string...")
	_, _ = fmt.Scan(&UserInput)

	// lets loop through the runes in the string
	var checks int
	for i, rune := range UserInput {
		if (rune == 'i' || rune == 'I') && i == 0 {
			checks++
		} else if rune == 'a' || rune == 'A' {
			checks++
		} else if (rune == 'n' || rune == 'N') && i == len(UserInput)-1 {
			checks++
		}

	}
	if checks >= 3 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
