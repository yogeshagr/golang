package main

import "fmt"

func print(a *int) {
	fmt.Println(*a)
}

func main() {
	var a = 10
	print(&a)
}
