// Dup2 prints the count and text of lines that appear more than once in the
// input. It reads from stdin or from a list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	foundIn := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, foundIn)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, foundIn)
			f.Close()
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\t%v\n", count, line, foundIn[line])
		}
	}
}


func isFileIncluded(file string, files []string) bool {
	for _, f := range files {
		if file == f {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, foundIn map[string][]string) {
	input := bufio.NewScanner(f)
	var line string
	for input.Scan() {
		line = input.Text()
		counts[line]++
		if (! isFileIncluded(f.Name(), foundIn[line])) {
			foundIn[line] = append(foundIn[line], f.Name())
		}
	}
}
