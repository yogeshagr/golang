// Dup3 prints count and text of lines that appear more than once in the named
// input files

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		countLines(data, counts)
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%v\n", count, line)
		}
	}
}

func countLines(data []byte, counts map[string]int) {
	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}
}
