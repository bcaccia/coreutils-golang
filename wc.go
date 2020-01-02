package main

import (
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
	"strings"
)

type results struct {
	Bytes   int64
	Chars   int64
	Lines   int64
	MaxLine int64
	Words   int64
}

// determines if args have been passed by checking the flag states in an
// array. if it encounters a true state, the loop breaks and returns true
func checkForFlags(flagStates []bool) (argsPassed bool) {
	for _, element := range flagStates {
		if element == true {
			argsPassed = true
			break
		}
	}
	return argsPassed
}

func printResults(flagsResult bool, flagStates []bool, r *results) {
	// if no flags have been passed print out the standard
	// newline, words, and byte counts
	if flagsResult == false {
		fmt.Println(r.Lines, r.Words, r.Bytes)
	}
}

func getCounts(r *results, scanner *bufio.Scanner) {
	// iterate through everything in the scanner and count the bytes
	for scanner.Scan() {
		// these two lines are for debugging only
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		//fmt.Println(int(len(scanner.Bytes())) + 1)

		// get the total bytes
		// add 1 for a human readable result
		r.Bytes += int64(len(scanner.Bytes()) + 1)

		// get the total characters
		// cast text as runes and get the len + 1
		// TODO getting a different char count on binary files when compared
		// to coreutils wc. suspect the different in count to be due to unicode
		r.Chars += int64(len([]rune(scanner.Text())) + 1)

		// get maximum line length
		// get the total chars but trim any whitespace off the
		// end. then compare if larger than previous result
		lenTotalTemp := int64(len([]rune(strings.TrimRight(scanner.Text(), " "))))
		if lenTotalTemp > r.MaxLine {
			r.MaxLine = lenTotalTemp
		}

		// get total words
		// break text up into words by white space characters
		// then calculate the length in the array
		r.Words += int64(len(strings.Fields(scanner.Text())))

		r.Lines++
	}
}

func main() {
	var r results
	// define all the available flags for this command
	byteFlag := flag.BoolP("bytes", "c", false, "print the byte counts")
	charsFlag := flag.BoolP("chars", "m", false, "print the character counts")
	linesFlag := flag.BoolP("lines", "l", false, "print the newline counts")
	maxLineLengthFlag := flag.BoolP("max-line-length", "L", false, "print the maximum display width")
	wordsFlag := flag.BoolP("words", "w", false, "print the maximum display width")
	versionFlag := flag.BoolP("version", "v", false, "output version information and exit")

	flag.Parse()

	// store all the flag states in an array
	var flagStates = []bool{*byteFlag, *charsFlag, *linesFlag, *maxLineLengthFlag, *wordsFlag, *versionFlag}
	// pass that array to a helper function to determine if args
	// have been passed
	flagsResult := checkForFlags(flagStates)

	// print out all the arguments passed in
	args := flag.Args()

	// check if any filename was passed as an arg. if so, take action
	if len(args) < 1 {
		// define a scanner to read from stdin
		scanner := bufio.NewScanner(os.Stdin)

		getCounts(&r, scanner)
		printResults(flagsResult, flagStates, &r)

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			// exit and indicate failure
			os.Exit(1)
		}
	} else {
		// iterate through all the args and perform actions
		for _, element := range args {
			file, err := os.Open(element)
			// define a scanner to read from the file
			scanner := bufio.NewScanner(file)

			getCounts(&r, scanner)
			printResults(flagsResult, flagStates, &r)

			if err != nil {
				fmt.Println("Failed to open the file: %s", element)
				// exit and indicate failure
				os.Exit(1)
			}
			defer file.Close()
		}
		// only print out the sum of all files if there is more
		// than one file in the args variable
		// TODO fix bug for lenResult. This should return the
		// longest line out of the two files, not sum them
		//if len(args) > 1 {
		//fmt.Println(bytesResult, charsResult, lenResult, wordsResult, linesResult, "total")
		//}

	}
}
