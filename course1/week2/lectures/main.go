/*
This is based on week 2
*/
package main

import (
	"fmt"
)

// pointer example -----------------------------------------------------
// https://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/
// golang by default is passed by value (opposite of reference, i.e a copy of argument is made)

type team struct {
	member1 string
	member2 string
	member3 string
}

func pointerFunc(t team) {
	// lets try to alter team object (struct)
	t.member1 = "germayne"

}

func pointerMain() {
	// lets create a team object
	t := team{
		member1: "rusty",
		member2: "john",
	}
	// pass into pointerFunc to alter
	pointerFunc(t)
	fmt.Println(t) // return {rusty john } no change at all
}

//--------------------------------------------------------------------
func pointerFunc2(t *team) {
	// lets try to alter team object (struct)
	t.member1 = "germayne"

}

func pointerMainCorrect() {
	// lets create a team object
	t := team{
		member1: "rusty",
		member2: "john",
	}
	// to alter the argument, we must use pointer
	pointerFunc2(&t) // we send the address
	fmt.Println(t)   // this will return {germayne john }
}

// anther pointer demo --------------------------------------

func pointerDemoAgain() {
	x := 1
	b := &x
	*b = 5
	fmt.Println(x)
}

// printf example --------------------------------------------------------
var x int = 1
var x2 float32 = 2.01221

func printfExampleMain() {
	fmt.Printf("variable x is %d \n", x)
	fmt.Printf("variable x is %f \n", x2)

}

// const example (constants) ----------------------------------------------

// func constDemo() {
// 	const x = 1.3
// 	x = 3
// 	fmt.Println(x) // get error: cannot assign to x
// }

// control flow -------------------------------------------------------------

// Scan demo ---------------------------------------------------------------

func scanner() {
	var userInput string
	fmt.Println("how old are you?")
	_, _ = fmt.Scan(&userInput) // Scan takes in pointer address and writes value to address
	fmt.Printf("You are %s \n", userInput)

}

// main function, must have -----------------------------------------------
func main() {

	//pointerMain()
	//pointerMainCorrect()
	//printfExampleMain()
	//constDemo()
	//pointerDemoAgain()
	scanner()
}
