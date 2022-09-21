package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// type ByAge []user
//
// func (a ByAge) Len() int           { return len(a) }
// func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByString []user

func (s ByString) Len() int           { return len(s) }
func (s ByString) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByString) Less(i, j int) bool { return s[i].First < s[j].First }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// sort.Sort(ByAge(users))
	// fmt.Println("--------- Sorting by Age ----------")
	// for _, v := range users {
	// 	fmt.Println(v.First, v.Last, v.Age)
	// 	for _, iv := range v.Sayings {
	// 		fmt.Println("\t\t", iv)
	// 	}
	// }
	fmt.Println("-------- Sorting by String -------")
	sort.Sort(ByString(users))
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		sort.Strings(v.Sayings)
		for _, iv := range v.Sayings {
			fmt.Println("\t\t", iv)
		}
	}
}
