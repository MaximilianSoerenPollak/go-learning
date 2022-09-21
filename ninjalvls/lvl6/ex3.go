package main

import "fmt"

func main() {

	fmt.Println("This is beforre the defer")
	defer (second)()
	fmt.Println("This is after the defer")

}

func second() {
	fmt.Println("I'm the second function")
}
