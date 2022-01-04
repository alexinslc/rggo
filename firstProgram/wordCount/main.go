package main

// test this program using the following:
// echo "My first command line tool with Go" | .\wordCount.exe

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")

	// Parsing the flags provided by the user
	flag.Parse()

	// Calling the count function to count the number of words (or lines)
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)
	// If the countlines flag and countbytes is not set, we want to count words so we
	// Define the scanner split type to words (default is split by lines)
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}
	// if -b is passed, count bytes
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	// Defining a counter
	wc := 0
	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}
	// Return the total
	return wc
}
