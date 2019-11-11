package main

import(
	"bufio"
	"fmt"
	"os"
)

func main(){
	files := os.Args[1:]
	counts := make(map[string]int)

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files{
			f, error := os.Open(arg)
			if error != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", error)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}


func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
