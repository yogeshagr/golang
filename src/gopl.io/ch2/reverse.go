package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return
}

func reverse2(a *[]int) {
	fmt.Println(len(a))
}

func main() {
	s := []int{1, 2, 3, 4, 5} // slice initialization
	reverse(s)

	a := [...]int{1, 2, 3, 4, 5} // array initialization
	reverse2(&a)

	reverse(a[:]) // passing a slice

	fmt.Println(s)
	fmt.Println(a)

	fmt.Printf("a is of type %T and s is of type %T\n", a, s)
}
