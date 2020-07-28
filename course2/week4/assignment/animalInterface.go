package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var name string
var info string

type Animal interface{
	Eat()
	Move()
	Speak()

}

type Cow struct {
	name string
}

func (c *Cow) Eat() {
	fmt.Println("grass")
}

func (c *Cow) Move() {
	fmt.Println("walk")
}

func (c *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	name string
}

func (b *Bird) Eat() {
	fmt.Println("worms")
}

func (b *Bird) Move() {
	fmt.Println("fly")
}

func (b *Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	name string
}

func (s *Snake) Eat() {
	fmt.Println("mice")
}

func (s *Snake) Move() {
	fmt.Println("slither")
}

func (s *Snake) Speak() {
	fmt.Println("hsss")
}


func main() {
	var exit bool

	for exit == false {
		fmt.Println("Enter Animal Name (e.g newanimal animaltype cow) : ")
		fmt.Println(">")
		var ani Animal

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		strArray := strings.Split(scanner.Text(), " ")

		// dynamic type and value
		// map concrete type to interface
		if strArray[2] == "cow"{
			cow := Cow{name: strArray[1]}
			ani = &cow //map Cow to animal so it Animal type
			fmt.Println("Created it!")
		} else if strArray[2] == "bird" {
			bird := Bird{name: strArray[1]}
			ani = &bird //map Bird to animal so it Animal type
			fmt.Println("Created it!")
		} else if strArray[2] == "snake" {
			snake := Snake{name: strArray[1]}
			ani = &snake //map Snake to animal so it Animal type
			fmt.Println("Created it!")
		}

		fmt.Println("Enter query (e.g query animalname animaltype) : ")
		fmt.Println(">")
		scanner2 := bufio.NewScanner(os.Stdin)
		scanner2.Scan()
		strArray2 := strings.Split(scanner2.Text(), " ")

		if strArray2[2] == "eat"{
			ani.Eat()
		} else if strArray2[2] == "move" {
			ani.Move()
		} else if strArray2[2] == "speak" {
			ani.Speak()

		}
	}
}