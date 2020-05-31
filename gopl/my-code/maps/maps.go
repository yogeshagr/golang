package main

import "fmt"

func a(k int, v int, x map[int]int) {
	x[k] = v
}

func main() {
	m := map[int]int{}
	k := 2
	v := 3
	a(k, v, m)
	fmt.Println(m)
}
