// Reports whether two strings are anagram of each other.
package main

import (
	"fmt"
	"os"
	"strings"
)

func isAnagram(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	if n1 == n2 {
		for i := 0; i < n1; i++ {
			if !strings.Contains(s1, string(s2[i])) || !strings.Contains(s2, string(s1[i])){
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func main(){
	s1 := os.Args[1]
	s2 := os.Args[2]

	fmt.Println(isAnagram(s1, s2))

}
