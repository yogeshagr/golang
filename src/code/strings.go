package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

  s1 := "left foot"
	fmt.Println(&s1)
	s1 += ", right foot"
	fmt.Println(&s1)
}
