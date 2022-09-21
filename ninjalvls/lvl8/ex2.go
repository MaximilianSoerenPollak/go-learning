package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":35}]`

	var people []person

	err := json.Unmarshal([]byte(s), &people) // Make sure you put a POINTER to the slice of struct
	if err != nil {                           // Not the struct itself. !!!
		fmt.Println(err)
	}

	fmt.Println(people[0].Age) // So here we have to access the slice first before we can access
	// the method. Important to remember.

}
