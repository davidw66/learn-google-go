package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pw := `p123`

	ep, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	//Why is has different each time?
	fmt.Println("Hashed password as a string -", string(ep))

	passwordEntered := "p123"

	err = bcrypt.CompareHashAndPassword(ep, []byte(passwordEntered))
	if err != nil {
		fmt.Println("CANNOT LOGIN", err)
		return
	}

	fmt.Println("Password matches - login approved")
}
