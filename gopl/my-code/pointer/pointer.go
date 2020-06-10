package main

import "fmt"

func main() {
	var a = 10
	var p *int
	p = &a
	fmt.Println(*p)
	fmt.Println(p)
	var q *int
	q = p
	fmt.Println(q)
	fmt.Println(p == q)
}
