// Wordfreq reports the frequency of each word in an input text file.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := os.Args[1]
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	input := bufio.NewScanner(f)

	if input.Err() != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}

	input.Split(bufio.ScanWords) // break the input into words instead of lines

	wordfreq := map[string]int{}

	for input.Scan() {
		word := input.Text()
		wordfreq[word]++
	}
	for k, v := range wordfreq {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
