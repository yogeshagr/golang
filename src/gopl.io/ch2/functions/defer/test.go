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

func triple3(x int) (result int) {
	test2 := func() {
		if result == 0 {
			result += x
		} else {
			result += 100
		}
	}
	defer test2()
	return x + x
}

// Eventhough yy is not defined inside the function body, yy is accessible
func triple4() (yy int) {
	yy = 2
	return yy
}

func main() {
	fmt.Println(triple1(4)) // "12"
	fmt.Println(triple2(4)) // "8"
	fmt.Println(triple3(4)) // "12"
	fmt.Println(triple4())  // "2"
}
