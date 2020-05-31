package main

import "fmt"

func modifySlice(a []string, b string) {
	a[0] = b
}

func main() {
	a := []string{"a", "b"}
	fmt.Println(a)
	modifySlice(a, "b")
	fmt.Println(a)
}
