package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func main() {
	p1 := person{}
	p2 := &person{name: "yogesh", age: 21}
	fmt.Println(p1, p1.name)
	fmt.Println(p2, p2.name)
}
