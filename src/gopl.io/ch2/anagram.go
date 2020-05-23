package main

import (
	"fmt"
	"os"
	"strings"
)

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if strings.IndexByte(s2, s1[i]) < 0 {
			return false
		}
	}
	return true
}

func main() {
	s1, s2 := os.Args[1], os.Args[2]
	fmt.Println(isAnagram(s1, s2))
}
