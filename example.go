package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) alter() string {
	r.width = r.width * 2
	return "success"
}

func (r *rect) alter2() string {
	r.width = r.width * 2
	return "success"
}

func main() {
	r := rect{width: 10, height: 5}
	r.alter()
	fmt.Println(r.width)

	r.alter2()
	fmt.Println(r.width)

}
