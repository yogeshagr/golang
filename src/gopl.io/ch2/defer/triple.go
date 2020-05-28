package main

import "fmt"

func triple1(x int) (result int) {
	defer func() { result += x }()
	return x + x
}

func triple2(x int) (result int) {
	defer test(x, result)
	return x + x
}

func test(x, result int) {
	result += x
}

func main() {
	fmt.Println(triple2(4)) // "8"
	fmt.Println(triple1(4)) // "12"
}
