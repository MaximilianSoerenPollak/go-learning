package main

import (
	"fmt"
)

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("My name is: ", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "Testy",
	}
	//saySomething(p1) Does not work because the interface of "speak"() is implemented as pointer (&person).
	saySomething(&p1) // Does work since we pass the pointer to person into the function (&p1)
	// p1.speak() This does work no matter which way it is implemented.
}
