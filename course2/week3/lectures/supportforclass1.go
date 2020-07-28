package main

import "fmt"

// MyInt -
type MyInt int

// Double - a method for MyInt object
func (mi MyInt) Double() int {
	return int(mi * 2)

}

func main() {
	v := MyInt(5)

	v.Double()

	// v is not changed because the method does a call by value (copy)
	fmt.Println(v)

}
