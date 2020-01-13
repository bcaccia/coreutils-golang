package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"os/user"
)

func main() {
	// define all the available flags for this command
	versionFlag := flag.BoolP("version", "v", false, "output version information and exit")
	flag.Parse()

	// check if the -v flag is passed and if so print out version information
	if *versionFlag == true {
		fmt.Println("whoami Golang rewrite v1.0")
		fmt.Println("Written by Benjamin Caccia")
	} else {

		// get current user data and store in variable
		user, err := user.Current()
		// handle potential errors by panicking
		if err != nil {
			panic(err)
		}

		// print out just the username
		fmt.Println(user.Name)
	}
}
