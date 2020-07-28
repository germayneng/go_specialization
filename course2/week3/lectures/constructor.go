/*
Constructor good practice
*/

package main

import "fmt"

type Person struct {
	name string
	job string
	nationality string
}

// here, the constructor will return a pointer
// for consistency since method receiver takes in pointer as well
func newPerson(name string,class string, nationality string ) *Person {
	return &Person{name: name,job: class, nationality: nationality}

}

// here we want our method to be able to "alter" the job field
// since constructor returns pointer, let us also make sure method receiver is pointer
func (p *Person) changeJob() {
	p.job = "new unknown job"
}

func main() {
	x := newPerson("germayne","developer","singapore")
	y := newPerson("xt","developer","singapore")

	fmt.Println(x.name)
	fmt.Println(y.name)
	fmt.Println(x.job)
	fmt.Println(y.job)
	fmt.Println("changing jobs-----")
	x.changeJob()
	y.changeJob()
	fmt.Println(x.name)
	fmt.Println(y.name)
	fmt.Println(x.job)
	fmt.Println(y.job)



}
