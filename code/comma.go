// Count three places to the left to place the first comma.
// Continue placing commas after every three digits.

package main

import "fmt"

func comma(s string) string {
	a := ""
	i := len(s)
	for ; i > 3; i -= 3 {
		a = "," + s[i-3:i] + a
	}
	a = s[:i] + a
	return a
}

func main(){
	fmt.Println("12345 =", comma("12345"))
	fmt.Println("1234567 =", comma("1234567"))
}
