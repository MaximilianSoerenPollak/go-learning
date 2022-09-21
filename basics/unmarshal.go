package main

import (
	"encoding/json"
	"fmt"
)

type person1 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"max","Last":"test","Age":28},{"First":"Tim","Last":"wowe","Age":30}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)
	var people []person1

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Println("\n Person Number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
