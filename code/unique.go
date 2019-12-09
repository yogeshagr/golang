package main

import "fmt"

func unique(a []string) []string {
	count := 0
	for _, s := range(a) {
		if a[count] == s {
			continue
		}
		count++
		a[count] = s
	}
	return a[:count+1]
}

func main() {
	a := []string{"a"}
	fmt.Println(a)
	fmt.Println(unique(a))
}
