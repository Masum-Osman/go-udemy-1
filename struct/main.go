package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Mocha",
	}

	alex.print()
	alex.updateName("Fright")
	alex.print()
}

func (p person) print() {
	fmt.Println(p)
}

func (p person) save() {
	fmt.Println(p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
