package main

import (
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

func main() {
	// define all the available flags for this command
	byteFlag := flag.BoolP("bytes", "c", false, "print the byte counts")
	charsFlag := flag.BoolP("chars", "m", false, "print the character counts")
	linesFlag := flag.BoolP("lines", "l", false, "print the newline counts")
	maxLineLengthFlag := flag.BoolP("max-line-length", "L", false, "print the maximum display width")
	wordsFlag := flag.BoolP("words", "w", false, "print the maximum display width")
	versionFlag := flag.BoolP("version", "v", false, "output version information and exit")

	flag.Parse()

	// print out flag states just so we can test
	fmt.Println(*byteFlag)
	fmt.Println(*charsFlag)
	fmt.Println(*linesFlag)
	fmt.Println(*maxLineLengthFlag)
	fmt.Println(*wordsFlag)
	fmt.Println(*versionFlag)
	// print out all the arguments passed in
	args := flag.Args()
	fmt.Println(args)

	// check if any filename was passed as an arg. if so, take action
	if len(args) < 1 {
		fmt.Println("no args, let's parse from stdin")
		// define a scanner to read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		// print out stdin for debugging
		var bytesTotal = 0
		for scanner.Scan() {
			//fmt.Println(scanner.Text()) // Println will add back the final '\n'
			fmt.Println(len(scanner.Bytes()))
			bytesTotal += len(scanner.Bytes())
		}
		fmt.Println(bytesTotal)
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	} else {
		fmt.Println("args provided, let's parse out the filename")
		fmt.Println(args)
	}

}
