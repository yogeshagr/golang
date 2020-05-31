package main

import "fmt"

func a(x int) {
	fmt.Println(x)
}

func a(s string) {
	fmt.Println(s)
}

func main() {
	a("abc")
	a(2)
}
