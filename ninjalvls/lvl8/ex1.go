package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	first string
	age   int
}

func main() {

	u1 := user{
		first: "James",
		age:   35,
	}
	u2 := user{
		first: "Enemy",
		age:   20,
	}
	users := []user{u1, u2}
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)

}
