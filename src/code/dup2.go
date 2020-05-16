package main

import(
	"bufio"
	"fmt"
	"os"
)

func main(){
	files := os.Args[1:]
	counts := make(map[string]int)
	fileNames := make(map[string][]string)

	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames)
	} else {
		for _, arg := range files{
			f, error := os.Open(arg)
			if error != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", error)
				continue
			}
			countLines(f, counts, fileNames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Println(fileNames[line][:])
		}
	}
}


func in(fileName string, strings []string) bool {
	for _, name := range strings {
		if fileName == name {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, fileNames map[string][]string){
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if !in(f.Name(), fileNames[input.Text()]) {
			fileNames[input.Text()] = append(fileNames[input.Text()], f.Name())
		}
	}
}
