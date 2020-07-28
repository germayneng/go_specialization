package main

import (
	"fmt"
	"math"
)

// GenDisplaceFn - a function that returns function
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {

	fn := func(t float64) float64 {
		output := (0.5 * a * math.Pow(t, 2)) + (v0 * t) + (s0)
		return output
	}
	return fn

}

func main() {
	var UserInputa float64
	var UserInputv float64
	var UserInputs float64
	var UserInputt float64

	fmt.Println("Please input an acceleration...")
	_, _ = fmt.Scan(&UserInputa)

	fmt.Println("Please input an velocity...")
	_, _ = fmt.Scan(&UserInputv)

	fmt.Println("Please input an displacement...")
	_, _ = fmt.Scan(&UserInputs)

	fmt.Println("Please input an time...")
	_, _ = fmt.Scan(&UserInputt)

	fn := GenDisplaceFn(UserInputa, UserInputv, UserInputs)

	// fn is the returned function which we can now just
	// pass in t
	fmt.Println(fn(UserInputt))

}
