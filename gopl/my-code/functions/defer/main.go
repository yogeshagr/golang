package main

import "fmt"

func deferFunction() {
	fmt.Println("inside defer statement-1")
	fmt.Println("inside defer statement-2")
	fmt.Println("inside defer statement-3")

}

func main() {
	fmt.Println("main started")
	defer deferFunction()
	fmt.Println("main ended")
}
