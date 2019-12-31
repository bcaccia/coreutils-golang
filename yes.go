package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

func main() {
	// define all the available flags for this command
	versionFlag := flag.BoolP("version", "v", false, "output version information and exit")

	flag.Parse()
	fmt.Println(*versionFlag)

	// get all the arguments passed to the command and store them in a variable
	args := flag.Args()

	// if no args are provided, just print out the default y with a newline
	if len(args) < 1 {
		for {
			fmt.Println("y")
		}
	} else {
		// infinite outer loop to repeatedly print the args
		// a newline will be printed after each iteration
		for {
			// the inner loop iterates through all the args and prints them on a single line
			for index, element := range args {
				// evaluate if this is the last arg or not
				// if it isn't, print the arg with an extra space at the end
				// if it is, just print the arg without the space
				if index != len(args)-1 {
					fmt.Print(element + " ")
				} else {
					fmt.Print(element)
				}
			}
			fmt.Print("\n")
		}
	}
}
