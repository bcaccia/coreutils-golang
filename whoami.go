package main

import (
	"fmt"
	"os/user"
)

func main() {
	// get current user data and store in variable
	user, err := user.Current()
	// handle potential errors by panicking
	if err != nil {
		panic(err)
	}

	// print out just the username
	fmt.Println(user.Name)
}
