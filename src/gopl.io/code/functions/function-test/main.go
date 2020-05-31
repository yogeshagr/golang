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

func test() {
	a := 10
	fmt.Println(a)

	{
		a := 20
		fmt.Println(a)
	}
}

func main() {
	fmt.Printf("%T\n%T\n", a, b)
	f := embedded()
	fmt.Println(f())
	fmt.Println(f())
	test()
}
