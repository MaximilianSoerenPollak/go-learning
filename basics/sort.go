package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{4, 5, 8, 1, 2, 10, 24, 1, 93, -12, 3}
	xs := []string{"Max", "James", "Test", "Anne", "Zena"}

	sort.Ints(xi) // Do not have to reasign since this function changes the underlying Array.
	fmt.Println(xi)
	// -------
	sort.Strings(xs)
	fmt.Println(xs)

	// Note: It seems that they only take one parameter. Therefore you can not tell it to sort the other
	//      way around. This is interesting, need to find ways how to do that in the future.
}
