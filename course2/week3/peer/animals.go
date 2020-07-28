package main

import (
	"fmt"
)

type animal struct {
	food, locomotion, sound string
}

func (v animal) eat() {
	fmt.Println(v.food)
}

func (v animal) move() {
	fmt.Println(v.locomotion)
}

func (v animal) speak() {
	fmt.Println(v.sound)
}

func main() {
	m := map[string]animal{
		"cow":   animal{"grass", "walk", "moo"},
		"bird":  animal{"worms", "fly", "peep"},
		"snake": animal{"mice", "slither", "hsss"},
	}
	for {
		fmt.Print("> Animal <enter> Action")
		animal := "0"
		action := "0"
		fmt.Scan(&animal, &action)
		if action == "eat" {
			m[animal].eat()
		} else if action == "move" {
			m[animal].move()
		} else if action == "speak" {
			m[animal].speak()
		}
	}
}
