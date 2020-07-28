package main

import (
	"fmt"
)

var name string
var info string

type Animal struct {
	food       string
	locomotive string
	noise      string
}

// Eat -
func (a *Animal) Eat() {
	if name == "cow" {
		a.food = "grass"
	} else if name == "bird" {
		a.food = "worms"
	} else if name == "snake" {
		a.food = "mice"
	}
}

// Move -
func (a *Animal) Move() {
	if name == "cow" {
		a.locomotive = "walk"
	} else if name == "bird" {
		a.locomotive = "fly"
	} else if name == "snake" {
		a.locomotive = "slither"
	}
}

// Speak -
func (a *Animal) Speak() {
	if name == "cow" {
		a.noise = "moo"
	} else if name == "bird" {
		a.noise = "peep"
	} else if name == "snake" {
		a.noise = "hsss"
	}
}

func main() {
	var exit bool

	for exit == false {

		fmt.Println("Enter Animal Name: ")
		fmt.Println(">")
		_, _ = fmt.Scan(&name)

		fmt.Println("Enter Animal Information: ")
		fmt.Println(">")
		_, _ = fmt.Scan(&info)

		if info == "eat" {
			animal := Animal{food: "empty", locomotive: "empty", noise: "empty"}
			animal.Eat()
			fmt.Println(animal.food)
		} else if info == "move" {
			animal := Animal{food: "empty", locomotive: "empty", noise: "empty"}
			animal.Move()
			fmt.Println(animal.locomotive)
		} else if info == "speak" {
			animal := Animal{food: "empty", locomotive: "empty", noise: "empty"}
			animal.Speak()
			fmt.Println(animal.noise)
		}

	}
}
