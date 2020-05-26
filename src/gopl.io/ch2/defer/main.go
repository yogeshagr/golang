package main

import "fmt"

func deferFunction() func() {
	fmt.Println("inside defer statement-1")
	fmt.Println("inside defer statement-2")
	fmt.Println("inside defer statement-3")
	return func() { fmt.Println("anonymous function") }
}

func main() {
	fmt.Println("main started")
	defer deferFunction()()
	fmt.Println("main ended")
}
