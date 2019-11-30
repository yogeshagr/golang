// Count three places to the left to place the first comma.
// Continue placing commas after every three digits.

package main

import "fmt"

func numeral(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return numeral(s[:n-3]) + "," + s[n-3:]
}

func main(){
	fmt.Println("12345 =", numeral("12345"))
	fmt.Println("1234567 =", numeral("1234567"))
	fmt.Println("1234 =", numeral("1234"))
	fmt.Println("123 =", numeral("123"))
}
