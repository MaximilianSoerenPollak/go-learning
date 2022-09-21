package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last)

}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   34,
	}
	p1.speak()
}
