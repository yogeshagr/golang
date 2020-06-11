package main

import "fmt"

func modifySlice(a []string, b string) {
	a[0] = b
}

func main() {
	a := []string{"a", "b"}
	fmt.Println(a) // [a b]
	modifySlice(a, "b")
	fmt.Println(a) // [b b]

	b := make([]int, 2, 4)
	b[0] = 1
	b[1] = 2
	b[3] = 3 // panic: runtime error: index out of range [3] with length 2
	// b = append(b, 3), this will work.
}
