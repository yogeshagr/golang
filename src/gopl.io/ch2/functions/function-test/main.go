package main

import "fmt"

func a(x int, y int) {
}

func b(y int, x int) {
}

func embedded() func() int {
	x := 0
	square := func() int {
		x++
		return x * x
	}
	return square
}

func main() {
	fmt.Printf("%T\n%T\n", a, b)
	f := embedded()
	fmt.Println(f())
	fmt.Println(f())
}
