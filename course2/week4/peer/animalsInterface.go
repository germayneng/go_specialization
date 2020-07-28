package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

type Bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

type Snake struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func main() {
	var command, name, data string
	mapa := make(map[string]Animal)
	fmt.Println("Type X to exit:")
	for {
		fmt.Printf(">")
		fmt.Scanf("%s %s %s", &command, &name, &data)
		if command == "X" {
			break
		}
		if command == "newanimal" {
			var animal Animal
			switch data {
			case "cow":
				animal = Cow{name, "grass", "walk", "moo"}
			case "bird":
				animal = Bird{name, "worms", "fly", "peep"}
			case "snake":
				animal = Snake{name, "mice", "slither", "hsss"}
			default:
				fmt.Println("Not valid animal")
				continue
			}
			mapa[name] = animal
			fmt.Printf("Created a %s named %s!\n", data, name)

		} else if command == "query" {
			animal := mapa[name]
			if animal == nil {
				fmt.Println("This animal does not exist")
				continue
			}
			switch data {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Not valid action")
				continue
			}
		} else {
			fmt.Println("Sorry, keyword not recognized.")
			continue
		}
	}
}
