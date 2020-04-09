package main

import (
	"fmt"
)

type contact struct {
	email   string
	contact int
}
type person struct {
	name    string
	age     int
	contact contact
}

func main() {
	a := person{name: "alok", age: 10, contact: contact{
		email:   "alok@gmail.com",
		contact: 8888777,
	},
	}
	// personPointers := &a
	// personPointers.updateName("avi")
	a.updateName("asvi")
	a.print()
}

func (p *person) updateName(name string) {
	(*p).name = name
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
