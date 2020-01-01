package main

import (
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

func getTotalBytes(scanner *bufio.Scanner) (total int) {

	// define a variable to hold the bytes total
	var bytesTotal = 0

	// iterate through everything in the scanner and count the bytes
	for scanner.Scan() {
		// these two lines are for debugging only
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		//fmt.Println(int(len(scanner.Bytes())) + 1)

		// convert the bytes on the line to an integer and add 1 to make it human readable
		bytesTotal += int(len(scanner.Bytes()) + 1)
	}
	// print out the computed total bytes
	return bytesTotal
}

func getTotalChars(scanner *bufio.Scanner) (total int) {

	// define a variable to hold the chars total
	var charsTotal = 0

	// iterate through everything in the scanner and count the chars
	for scanner.Scan() {
		// convert and store contents of scanner to text
		scannerText := scanner.Text()
		// cast text as runes and get the len + 1
		// TODO getting a different char count on binary files when compared
		// to coreutils wc. suspect the different in count to be due to unicode
		charsTotal += len([]rune(scannerText)) + 1
	}
	// print out the computed total chars
	return charsTotal
}

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

	// check if any filename was passed as an arg. if so, take action
	if len(args) < 1 {
		fmt.Println("no args, let's parse from stdin")
		// define a scanner to read from stdin
		scanner := bufio.NewScanner(os.Stdin)

		//total := getTotalBytes(scanner)
		total := getTotalChars(scanner)

		fmt.Println(total)
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			// exit and indicate failure
			os.Exit(1)
		}
	} else {
		fmt.Println("args provided, let's parse out the filename")
		fmt.Println(args)
		// iterate through all the args and perform actions
		for _, element := range args {
			file, err := os.Open(element)
			// define a scanner to read from the file
			scanner := bufio.NewScanner(file)

			//total := getTotalBytes(scanner)
			total := getTotalChars(scanner)

			fmt.Println(total)

			if err != nil {
				fmt.Println("Failed to open the file: %s", element)
				// exit and indicate failure
				os.Exit(1)
			}
			defer file.Close()
		}

	}
}
