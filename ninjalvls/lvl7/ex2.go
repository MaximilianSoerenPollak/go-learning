package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person, n int) {
	p.age = n
}

func main() {
	p1 := person{
		first: "Max",
		last:  "Pol",
		age:   28,
	}
	fmt.Println(p1)
	changeMe(&p1, 30)
	fmt.Println(p1)
}
