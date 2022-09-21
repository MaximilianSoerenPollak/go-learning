package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	s := `verystrongpw`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	loginPassword1 := `verystrongpw1`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword1))
	if err != nil {
		fmt.Println(err)
		fmt.Println("You can not log in")

	}
	fmt.Println("You're logged in")
}
