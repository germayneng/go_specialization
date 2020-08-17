package main

import (
	"fmt"
)

func addX(x *int) {
	for i := 0; i < 4; i++ {
		(*x)++

	}
}
func main() {
	x := 100
	go addX(&x)
	for i := 0; i < 4; i++ {
		fmt.Println(x)
	}
}
