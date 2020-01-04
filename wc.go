package main

import (
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
	"strings"
)

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

func printResults(flagsResult bool, flagStates []bool, resultsArray []uint64) {
	// print out default values if false
	// these are newline, words, and byte counts
	if flagsResult == false {
		fmt.Print(resultsArray[2], resultsArray[4], resultsArray[0])
		fmt.Print(" ")
	} else {
		// The counts are printed in this order: newlines, words, characters, bytes, maximum line length
		for index, element := range resultsArray {
			if flagStates[index] == true {
				fmt.Print(element)
				fmt.Print(" ")
			}
		}
	}

}

func getCounts(scanner *bufio.Scanner) (result []uint64) {

	// define a variable to hold the bytes total
	var linesTotal uint64
	var wordsTotal uint64
	var charsTotal uint64
	var bytesTotal uint64
	var lenTotal uint64

	var resultsArray []uint64

	// iterate through everything in the scanner and count the bytes
	for scanner.Scan() {
		// these two lines are for debugging only
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		//fmt.Println(int(len(scanner.Bytes())) + 1)

		// get the total bytes
		// add 1 for a human readable result
		bytesTotal += uint64(len(scanner.Bytes()) + 1)

		// get the total characters
		// cast text as runes and get the len + 1
		// TODO getting a different char count on binary files when compared
		// to coreutils wc. suspect the different in count to be due to unicode
		charsTotal += uint64(len([]rune(scanner.Text())) + 1)

		// get maximum line length
		// get the total chars but trim any whitespace off the
		// end. then compare if larger than previous result
		lenTotalTemp := uint64(len([]rune(strings.TrimRight(scanner.Text(), " "))))
		if lenTotalTemp > lenTotal {
			lenTotal = lenTotalTemp
		}

		// get total words
		// break text up into words by white space characters
		// then calculate the length in the array
		wordsTotal += uint64(len(strings.Fields(scanner.Text())))

		linesTotal++
	}
	// return the computed total bytes
	resultsArray = append(resultsArray, linesTotal, wordsTotal, charsTotal, bytesTotal, lenTotal)
	return resultsArray

}

func main() {
	var filesFrom string

	// define all the available flags for this command
	byteFlag := flag.BoolP("bytes", "c", false, "print the byte counts")
	charsFlag := flag.BoolP("chars", "m", false, "print the character counts")
	linesFlag := flag.BoolP("lines", "l", false, "print the newline counts")
	flag.StringVar(&filesFrom, "files0-from", "", `read input from the files specified by
                             NUL-terminated names in file F;
                             If F is - then read names from standard input`)
	maxLineLengthFlag := flag.BoolP("max-line-length", "L", false, "print the maximum display width")
	wordsFlag := flag.BoolP("words", "w", false, "print the maximum display width")
	versionFlag := flag.BoolP("version", "v", false, "output version information and exit")

	flag.Parse()

	// store all the flag states in an array
	var flagStates = []bool{*linesFlag, *wordsFlag, *charsFlag, *byteFlag, *maxLineLengthFlag}
	// pass that array to a helper function to determine if args
	// have been passed
	flagsResult := checkForFlags(flagStates)

	// print out all the arguments passed in
	args := flag.Args()

	// check if the -v flag is passed and if so print out version information
	if *versionFlag == true {
		fmt.Println("wc Golang rewrite v1.0")
		fmt.Println("Written by Benjamin Caccia")
	} else if len(filesFrom) > 0 {
		fmt.Println(filesFrom)
		fmt.Println("not yet implemented")

	} else {

		// check if any filename was passed as an arg. if so, take action
		if len(args) < 1 {
			// define a scanner to read from stdin
			scanner := bufio.NewScanner(os.Stdin)

			//bytesResult, charsResult, lenResult, wordsResult, linesResult := getCounts(scanner)
			var resultsArray []uint64
			resultsArray = getCounts(scanner)
			printResults(flagsResult, flagStates, resultsArray)

			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
				// exit and indicate failure
				os.Exit(1)
			}
		} else {
			// declare vars used to tally the results for each file
			//var bytesResult, charsResult, lenResult, wordsResult, linesResult int
			var totalTally [5]uint64
			// iterate through all the args and perform actions
			for _, element := range args {
				file, err := os.Open(element)
				// define a scanner to read from the file
				scanner := bufio.NewScanner(file)

				//bytesResultTemp, charsResultTemp, lenResultTemp, wordsResultTemp, linesResultTemp := getCounts(scanner)
				var resultsArray []uint64
				resultsArray = getCounts(scanner)
				printResults(flagsResult, flagStates, resultsArray)
				fmt.Print(element + "\n")

				//fmt.Println(bytesResultTemp, charsResultTemp, lenResultTemp, wordsResultTemp, linesResultTemp, element)

				for index, element := range resultsArray {
					totalTally[index] += element
				}

				// add results to the tally variables
				//bytesResult += bytesResultTemp
				//charsResult += charsResultTemp
				//lenResult += lenResultTemp
				//wordsResult += wordsResultTemp
				//linesResult += linesResultTemp

				if err != nil {
					fmt.Println("Failed to open the file: %s", element)
					// exit and indicate failure
					os.Exit(1)
				}
				defer file.Close()
			}
			// only print out the sum of all files if there is more
			// than one file in the args variable
			if len(args) > 1 {

				if flagsResult == false {
					fmt.Print(totalTally[2], totalTally[4], totalTally[0])
					fmt.Print(" ")
				} else {
					for index, element := range totalTally {
						if flagStates[index] == true {
							fmt.Print(element)
							fmt.Print(" ")
						}
					}
				}
				fmt.Print("total")
			}

		}
	}
}
