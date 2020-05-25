package main

import "fmt"

func a(x int, y int) {
}

func b(y int, x int) {
}

func main() {
	fmt.Printf("%T\n%T\n", a, b)
}
