package main

import "fmt"

//Demo1 a function that demostrates how to do a "reference" style in golang--------------
// we are passing pointer object to the function
func alterInt(y *int) {
	*y = *y + 1
}

func mainAlterInt() {

	//create an int
	x := 0
	alterInt(&x)
	fmt.Println(x)

}

//Demo2 a function that demostrates a reference style in golang for struc-------------
// we are passing pointer object to the function
type Person struct {
	Name string
}

func alterStruct(y *Person) {
	y.Name = "germayne"

}

func main() {
	// for demo 1
	//mainAlterInt()

	// for demo 2
	// x := Person{
	// 	Name: "rusty"}
	// fmt.Println(x.Name)
	// alterStruct(&x)
	// fmt.Println(x.Name)

	// demo 3
	x := Person{Name: "rus"}

	var y *Person
	y = &x
	y.Name = "germ"
	fmt.Println(x)

}
